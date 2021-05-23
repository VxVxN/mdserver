package post

import (
	"net/http"

	"github.com/VxVxN/log"
)

func (ctrl *Controller) EditPostHandler(w http.ResponseWriter, r *http.Request) {
	if errObj := ctrl.getPost(w, r, true); errObj != nil {
		log.Error.Printf("Failed to edit post: %v", errObj.Error)
		errObj.JsonResponse(w)
		return
	}
}
