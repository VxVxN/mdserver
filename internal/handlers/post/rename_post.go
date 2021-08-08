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

type RequestRenamePost struct {
	DirName     string `json:"dir_name" binding:"required"`
	OldFileName string `json:"old_file_name" binding:"required"`
	NewFileName string `json:"new_file_name" binding:"required"`
}

func (ctrl *Controller) RenamePostHandler(c *gin.Context) {
	var req RequestRenamePost

	errObj := tools.UnmarshalRequest(c, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	if errObj = ctrl.renamePost(c, req); errObj != nil {
		log.Error.Printf("Failed to rename post: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) renamePost(c *gin.Context, req RequestRenamePost) *e.ErrObject {
	username, err := tools.GetUserNameFromSession(c)
	if err != nil {
		err = fmt.Errorf("cannot get username from session: %v", err)
		return e.NewError("Failed to get username from session", http.StatusBadRequest, err)
	}

	oldPathToDir := path.Join(consts.PathToPosts, username, req.DirName, req.OldFileName)
	oldPathToDir += consts.ExtMd

	newPathToDir := path.Join(consts.PathToPosts, username, req.DirName, req.NewFileName)
	newPathToDir += consts.ExtMd

	if errObj := ctrl.mongoPosts.RenamePost(username, req.DirName, req.OldFileName, req.NewFileName); errObj != nil {
		return errObj
	}

	if err = os.Rename(oldPathToDir, newPathToDir); err != nil {
		err = fmt.Errorf("can't rename post: %v", err)
		return e.NewError("Failed to rename post", http.StatusInternalServerError, err)
	}

	oldPrefix := req.DirName + "_" + req.OldFileName
	newPrefix := req.DirName + "_" + req.NewFileName
	if err = renameImagesByPrefix(username, oldPrefix, newPrefix); err != nil {
		err = fmt.Errorf("can't delete images bind with file: %v", err)
		return e.NewError("Failed to delete images bind with file", http.StatusInternalServerError, err)
	}

	return nil
}
