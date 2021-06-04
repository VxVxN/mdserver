package post

import (
	"fmt"
	"net/http"

	e "github.com/VxVxN/mdserver/pkg/error"

	"github.com/VxVxN/log"
)

func (ctrl *Controller) EditPostHandler(w http.ResponseWriter, r *http.Request) {
	if errObj := ctrl.getEditingPost(w, r); errObj != nil {
		log.Error.Printf("Failed to edit post: %v", errObj.Error)
		errObj.JsonResponse(w)
		return
	}
}

func (ctrl *Controller) getEditingPost(w http.ResponseWriter, r *http.Request) *e.ErrObject {
	postMD := ctrl.getPathToPostMD(r)

	templatePost, status, err := ctrl.posts.Get(postMD, true)
	if err != nil {
		err = fmt.Errorf("can't get post: %v, post: %s", err, postMD)
		return e.NewError("Failed to get post", status, err)
	}

	if err = ctrl.editingPostTemplate.ExecuteTemplate(w, "layout", templatePost); err != nil {
		err = fmt.Errorf("can't execute template: %v, post: %s", err, templatePost.Title)
		return e.NewError("Failed to get post", http.StatusInternalServerError, err)
	}
	return nil
}
