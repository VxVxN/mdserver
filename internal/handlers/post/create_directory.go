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

type RequestCreateDirectory struct {
	DirName string `json:"name" binding:"required"`
}

/**
 * @api {post} /create_directory Create a directory
 * @apiName CreateDirectoryHandler
 * @apiGroup post
 *
 * @apiParamExample {json} Request example:
 * {
 *    "name":"Programming"
 * }
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 500 Internal Server Error
 * {
 *    "message":"Failed to create directory"
 * }
 */

func (ctrl *Controller) CreateDirectoryHandler(c *gin.Context) {
	var req RequestCreateDirectory

	errObj := tools.UnmarshalRequest(c, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	if errObj = ctrl.createDirectory(c, req); errObj != nil {
		log.Error.Printf("Failed to create directory: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) createDirectory(c *gin.Context, req RequestCreateDirectory) *e.ErrObject {
	username, err := tools.GetUserNameFromSession(c)
	if err != nil {
		err = fmt.Errorf("cannot get username from session: %v", err)
		return e.NewError("Failed to get username from session", http.StatusBadRequest, err)
	}

	pathToDir := path.Join(consts.PathToPosts, username, req.DirName)

	if err := os.Mkdir(pathToDir, 0777); err != nil {
		err = fmt.Errorf("can't create directory: %v", err)
		return e.NewError("Failed to create directory", http.StatusInternalServerError, err)
	}

	if errObj := ctrl.mongoPosts.CreateDirectory(username, req.DirName); errObj != nil {
		return errObj
	}
	return nil
}
