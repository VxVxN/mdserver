package post

import (
	"net/http"
)

func (ctrl *Controller) EditPostHandler(w http.ResponseWriter, r *http.Request) {
	ctrl.post(w, r, true)
}
