package post

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/VxVxN/log"
	"github.com/VxVxN/mdserver/pkg/consts"
	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/VxVxN/mdserver/pkg/tools"
)

type RequestDeleteDirectory struct {
	DirName string `json:"name"`
}

func (ctrl *Controller) DeleteDirectoryHandler(w http.ResponseWriter, r *http.Request) {
	var req RequestDeleteDirectory

	errObj := tools.UnmarshalRequest(r, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(w)
		return
	}

	if errObj = ctrl.deleteDirectory(req); errObj != nil {
		log.Error.Printf("Failed to delete directory: %v", errObj.Error)
		errObj.JsonResponse(w)
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
