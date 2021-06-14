package post

import (
	"fmt"
	"net/http"

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

	if err = ctrl.editingPostTemplate.ExecuteTemplate(c.Writer, "layout", templatePost); err != nil {
		err = fmt.Errorf("can't execute template: %v, post: %s", err, templatePost.Title)
		return e.NewError("Failed to get post", http.StatusInternalServerError, err)
	}
	return nil
}
