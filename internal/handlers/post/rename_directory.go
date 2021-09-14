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

	if errObj = ctrl.renameDirectory(c, req); errObj != nil {
		log.Error.Printf("Failed to rename directory: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) renameDirectory(c *gin.Context, req RequestRenameDirectory) *e.ErrObject {
	username, err := tools.GetUserNameFromSession(c)
	if err != nil {
		err = fmt.Errorf("cannot get username from session: %v", err)
		return e.NewError("Failed to get username from session", http.StatusBadRequest, err)
	}

	oldPathToDir := path.Join(consts.PathToPosts, username, req.OldDirName)
	newPathToDir := path.Join(consts.PathToPosts, username, req.NewDirName)

	if errObj := ctrl.mongoPosts.RenameDirectory(req.OldDirName, req.NewDirName); errObj != nil {
		return errObj
	}

	if err = os.Rename(oldPathToDir, newPathToDir); err != nil {
		err = fmt.Errorf("can't rename directory: %v", err)
		return e.NewError("Failed to rename directory", http.StatusInternalServerError, err)
	}

	return nil
}
