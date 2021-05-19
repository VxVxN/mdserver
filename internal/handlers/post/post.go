package post

import (
	"log"
	"mdserver/internal/glob"
	e "mdserver/pkg/error"
	"net/http"
	"path"
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
		postMD += "/index.md"
	}

	post, status, err := ctrl.posts.Get(postMD, isEdit)
	if err != nil {
		e.ErrorHandler(w, r, status)
		return
	}
	if err := ctrl.postTemplate.ExecuteTemplate(w, "layout", post); err != nil {
		log.Println(err.Error())
		e.ErrorHandler(w, r, 500)
	}
}
