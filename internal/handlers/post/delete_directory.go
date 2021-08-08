package post

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/log"

	"github.com/VxVxN/mdserver/internal/driver/mongo/posts"
	"github.com/VxVxN/mdserver/pkg/consts"
	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/VxVxN/mdserver/pkg/tools"
)

type RequestDeleteDirectory struct {
	DirName string `json:"name" binding:"required"`
}

func (ctrl *Controller) DeleteDirectoryHandler(c *gin.Context) {
	var req RequestDeleteDirectory

	errObj := tools.UnmarshalRequest(c, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	if errObj = ctrl.deleteDirectory(c, req); errObj != nil {
		log.Error.Printf("Failed to delete directory: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) deleteDirectory(c *gin.Context, req RequestDeleteDirectory) *e.ErrObject {
	username, err := tools.GetUserNameFromSession(c)
	if err != nil {
		err = fmt.Errorf("cannot get username from session: %v", err)
		return e.NewError("Failed to get username from session", http.StatusBadRequest, err)
	}

	pathToDir := path.Join(consts.PathToPosts, username, req.DirName)

	dir, err := ctrl.mongoPosts.GetDirectory(username, req.DirName)
	if err != nil {
		err = fmt.Errorf("cannot get directory: %v", err)
		return e.NewError("Failed to get directory", http.StatusInternalServerError, err)
	}

	if errObj := ctrl.mongoPosts.DeleteDirectory(req.DirName); errObj != nil {
		return errObj
	}

	if err = os.RemoveAll(pathToDir); err != nil {
		err = fmt.Errorf("can't delete directory: %v", err)
		return e.NewError("Failed to delete directory", http.StatusInternalServerError, err)
	}

	if err = removeImages(username, dir); err != nil {
		err = fmt.Errorf("can't delete images bind with dir: %v", err)
		return e.NewError("Failed to delete images bind with dir", http.StatusInternalServerError, err)
	}

	return nil
}

func removeImages(username string, dir *posts.Directory) error {
	var err error
	pathToImages := path.Join(consts.PathToPosts, username, "images")
	for _, file := range dir.Files {
		for _, image := range file.Images {
			pathToImage := path.Join(pathToImages, image.UUID)
			if err = os.Remove(pathToImage); err != nil {
				log.Warning.Printf("Failed to remove image: %v", err)
				return err
			}
		}
	}
	return nil
}
