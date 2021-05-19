package post

import (
	"log"
	"net/http"
	"path"
)

func (ctrl *Controller) PostHandler(w http.ResponseWriter, r *http.Request) {
	ctrl.post(w, r, false)
}

func (ctrl *Controller) post(w http.ResponseWriter, r *http.Request, isEdit bool) {
	params := r.URL.Query()

	page := params.Get(":page")
	postMD := path.Join(ctrl.workDir, "posts", page)

	if page == "" {
		// если page пусто, то выдаем главную
		postMD += "/index.md"
	}

	post, status, err := ctrl.posts.Get(postMD, isEdit)
	if err != nil {
		ctrl.errorHandler(w, r, status)
		return
	}
	if err := ctrl.postTemplate.ExecuteTemplate(w, "layout", post); err != nil {
		log.Println(err.Error())
		ctrl.errorHandler(w, r, 500)
	}
}

func (ctrl *Controller) errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	log.Printf("error %d %s %s\n", status, r.RemoteAddr, r.URL.Path)
	w.WriteHeader(status)
	if err := ctrl.errorTemplate.ExecuteTemplate(w, "layout", map[string]interface{}{"Error": http.StatusText(status), "Status": status}); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
}
