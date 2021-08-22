package post

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/log"

	"github.com/VxVxN/mdserver/pkg/consts"
	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/VxVxN/mdserver/pkg/tools"
)

type RequestDeletePost struct {
	DirName  string `json:"dir_name" binding:"required"`
	FileName string `json:"file_name" binding:"required"`
}

func (ctrl *Controller) DeletePostHandler(c *gin.Context) {
	var req RequestDeletePost

	errObj := tools.UnmarshalRequest(c, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	if errObj = ctrl.deletePost(c, req); errObj != nil {
		log.Error.Printf("Failed to delete post: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) deletePost(c *gin.Context, req RequestDeletePost) *e.ErrObject {
	username, err := tools.GetUserNameFromSession(c)
	if err != nil {
		err = fmt.Errorf("cannot get username from session: %v", err)
		return e.NewError("Failed to get username from session", http.StatusBadRequest, err)
	}

	pathToDir := path.Join(consts.PathToPosts, username, req.DirName, req.FileName)
	pathToDir += consts.ExtMd

	if err = ctrl.removeImages(username, req.DirName, req.FileName); err != nil {
		return e.NewError("Failed to delete images bind with post", http.StatusInternalServerError, err)
	}

	if errObj := ctrl.mongoPosts.DeletePost(req.DirName, req.FileName); errObj != nil {
		return errObj
	}

	if err = os.Remove(pathToDir); err != nil {
		err = fmt.Errorf("can't delete post: %v", err)
		return e.NewError("Failed to delete post", http.StatusInternalServerError, err)
	}

	return nil
}

func (ctrl *Controller) removeImages(username, dirName, fileName string) error {
	images, err := ctrl.mongoPosts.GetImages(username, dirName, fileName)
	if err != nil {
		return fmt.Errorf("cannot get images from mongo: %v", err)
	}
	for _, image := range images {
		pathToDestination := path.Join(consts.PathToImages, image.UUID)
		if err = os.Remove(pathToDestination); err != nil {
			return fmt.Errorf("cannot remove image: %v", err)
		}
	}
	return nil
}
