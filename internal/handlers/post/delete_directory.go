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

	if errObj = ctrl.deleteDirectory(req); errObj != nil {
		log.Error.Printf("Failed to delete directory: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) deleteDirectory(req RequestDeleteDirectory) *e.ErrObject {
	pathToDir := path.Join(consts.PathToPosts, req.DirName)

	if err := os.RemoveAll(pathToDir); err != nil {
		err = fmt.Errorf("can't delete directory: %v", err)
		return e.NewError("Failed to delete directory", http.StatusInternalServerError, err)
	}

	if errObj := ctrl.mongoPosts.DeleteDirectory(req.DirName); errObj != nil {
		return errObj
	}
	return nil
}
