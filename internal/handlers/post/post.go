package post

import (
	"log"
	"net/http"
	"path"

	"github.com/VxVxN/mdserver/internal/glob"
	"github.com/VxVxN/mdserver/pkg/consts"
)

func (ctrl *Controller) PostHandler(w http.ResponseWriter, r *http.Request) {
	ctrl.post(w, r, false)
}

func (ctrl *Controller) post(w http.ResponseWriter, r *http.Request, isEdit bool) {
	params := r.URL.Query()

	page := params.Get(":page")
	postMD := path.Join(glob.WorkDir, "posts", page)

	if page == "" {
		// если page пусто, то выдаем главную
		postMD += "/index"
	}
	postMD += consts.ExtMd

	post, status, err := ctrl.posts.Get(postMD, isEdit)
	if err != nil {
		ctrl.ErrorResponse(w, r, status)
		return
	}
	if err := ctrl.postTemplate.ExecuteTemplate(w, "layout", post); err != nil {
		log.Println(err.Error())
		ctrl.ErrorResponse(w, r, http.StatusInternalServerError)
	}
}
