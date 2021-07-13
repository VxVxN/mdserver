package post

import (
	"fmt"

	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/gin-gonic/gin"

	"github.com/VxVxN/log"
)

func (ctrl *Controller) EditPostHandler(c *gin.Context) {
	if errObj := ctrl.getEditingPost(c); errObj != nil {
		log.Error.Printf("Failed to edit post: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) getEditingPost(c *gin.Context) *e.ErrObject {
	postMD := ctrl.getPathToPostMD(c)

	templatePost, status, err := ctrl.posts.Get(postMD, true)
	if err != nil {
		err = fmt.Errorf("can't get post: %v, post: %s", err, postMD)
		return e.NewError("Failed to get post", status, err)
	}

	c.HTML(200, "editing_post.tmpl", templatePost)

	return nil
}
