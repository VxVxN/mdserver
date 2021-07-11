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

	if errObj = ctrl.deletePost(req); errObj != nil {
		log.Error.Printf("Failed to delete post: %v", errObj.Error)
		errObj.JsonResponse(c)
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
