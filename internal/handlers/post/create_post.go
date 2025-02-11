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

type RequestCreatePost struct {
	DirName  string `json:"dir_name" binding:"required"`
	FileName string `json:"file_name" binding:"required"`
}

/**
 * @api {post} /create_post Create a post
 * @apiName CreatePostHandler
 * @apiGroup post
 *
 * @apiParamExample {json} Request example:
 * {
 *    "dir_name":"Programming",
 *    "file_name":"Golang"
 * }
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 500 Internal Server Error
 * {
 *    "message":"Failed to create post"
 * }
 */

func (ctrl *Controller) CreatePostHandler(c *gin.Context) {
	var req RequestCreatePost

	errObj := tools.UnmarshalRequest(c, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	if errObj = ctrl.createPost(c, req); errObj != nil {
		log.Error.Printf("Failed to create post: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) createPost(c *gin.Context, req RequestCreatePost) *e.ErrObject {
	username, err := tools.GetUserNameFromSession(c)
	if err != nil {
		err = fmt.Errorf("cannot get username from session: %v", err)
		return e.NewError("Failed to get username from session", http.StatusBadRequest, err)
	}

	pathToDir := path.Join(consts.PathToPosts, username, req.DirName, req.FileName)
	pathToDir += consts.ExtMd

	file, err := os.Create(pathToDir)
	if err != nil {
		err = fmt.Errorf("can't create post: %v", err)
		return e.NewError("Failed to create post", http.StatusInternalServerError, err)
	}
	if err = file.Close(); err != nil {
		err = fmt.Errorf("can't close file: %v", err)
		return e.NewError("Failed to close file", http.StatusInternalServerError, err)
	}

	if errObj := ctrl.mongoPosts.CreatePost(username, req.DirName, req.FileName); errObj != nil {
		return errObj
	}
	return nil
}
