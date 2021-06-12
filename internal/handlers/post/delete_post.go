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

type RequestDeletePost struct {
	DirName  string `json:"dir_name"`
	FileName string `json:"file_name"`
}

func (ctrl *Controller) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	var req RequestDeletePost

	errObj := tools.UnmarshalRequest(r, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(w)
		return
	}

	if errObj = ctrl.deletePost(req); errObj != nil {
		log.Error.Printf("Failed to delete post: %v", errObj.Error)
		errObj.JsonResponse(w)
		return
	}
}

func (ctrl *Controller) deletePost(req RequestDeletePost) *e.ErrObject {
	pathToDir := path.Join(consts.PathToPosts, req.DirName, req.FileName)
	pathToDir += consts.ExtMd

	if err := os.Remove(pathToDir); err != nil {
		err = fmt.Errorf("can't delete post: %v", err)
		return e.NewError("Failed to delete post", http.StatusInternalServerError, err)
	}

	if errObj := ctrl.mongoPosts.DeletePost(req.DirName, req.FileName); errObj != nil {
		return errObj
	}
	return nil
}
