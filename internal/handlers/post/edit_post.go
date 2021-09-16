package post

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	e "github.com/VxVxN/mdserver/pkg/error"

	"github.com/VxVxN/log"
)

/**
 * @api {get} /edit/:dir/:file Get editing post page
 * @apiName EditPostHandler
 * @apiGroup post
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 500 Internal Server Error
 * {
 *    "message":"Failed to get post"
 * }
 */

func (ctrl *Controller) EditPostHandler(c *gin.Context) {
	if errObj := ctrl.getEditingPost(c); errObj != nil {
		log.Error.Printf("Failed to edit post: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) getEditingPost(c *gin.Context) *e.ErrObject {
	postMD, err := ctrl.getPathToPostMD(c)
	if err != nil {
		err = fmt.Errorf("can't get path to post: %v", err)
		return e.NewError("Failed to get path to post", http.StatusInternalServerError, err)
	}

	templatePost, status, err := ctrl.posts.Get(postMD, true)
	if err != nil {
		err = fmt.Errorf("can't get post: %v, post: %s", err, postMD)
		return e.NewError("Failed to get post", status, err)
	}

	c.HTML(200, "editing_post.tmpl", templatePost)

	return nil
}
