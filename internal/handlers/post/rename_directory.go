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

type RequestRenameDirectory struct {
	OldDirName string `json:"old_name" binding:"required"`
	NewDirName string `json:"new_name" binding:"required"`
}

func (ctrl *Controller) RenameDirectoryHandler(c *gin.Context) {
	var req RequestRenameDirectory

	errObj := tools.UnmarshalRequest(c, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	if errObj = ctrl.renameDirectory(req); errObj != nil {
		log.Error.Printf("Failed to rename directory: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) renameDirectory(req RequestRenameDirectory) *e.ErrObject {
	oldPathToDir := path.Join(consts.PathToPosts, req.OldDirName)
	newPathToDir := path.Join(consts.PathToPosts, req.NewDirName)

	if err := os.Rename(oldPathToDir, newPathToDir); err != nil {
		err = fmt.Errorf("can't rename directory: %v", err)
		return e.NewError("Failed to rename directory", http.StatusInternalServerError, err)
	}

	if errObj := ctrl.mongoPosts.RenameDirectory(req.OldDirName, req.NewDirName); errObj != nil {
		return errObj
	}
	return nil
}
