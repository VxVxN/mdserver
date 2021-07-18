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

	if errObj = ctrl.renamePost(req); errObj != nil {
		log.Error.Printf("Failed to rename post: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) renamePost(req RequestRenamePost) *e.ErrObject {
	oldPathToDir := path.Join(consts.PathToPosts, req.DirName, req.OldFileName)
	oldPathToDir += consts.ExtMd

	newPathToDir := path.Join(consts.PathToPosts, req.DirName, req.NewFileName)
	newPathToDir += consts.ExtMd

	err := os.Rename(oldPathToDir, newPathToDir)
	if err != nil {
		err = fmt.Errorf("can't rename post: %v", err)
		return e.NewError("Failed to rename post", http.StatusInternalServerError, err)
	}

	if errObj := ctrl.mongoPosts.RenamePost(req.DirName, req.OldFileName, req.NewFileName); errObj != nil {
		return errObj
	}
	return nil
}
