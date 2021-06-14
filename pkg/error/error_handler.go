package error

import (
	"html/template"
	"net/http"
	"path"

	"github.com/VxVxN/mdserver/pkg/consts"
)

type ErrResponseController struct {
	errorTemplate *template.Template
}

func (e *ErrResponseController) InitErrResponseController() {
	e.errorTemplate = template.Must(template.ParseFiles(path.Join(consts.PathToTemplates, "layout.html"), path.Join(consts.PathToTemplates, "error.html")))
}

func (e *ErrResponseController) ErrorResponse(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if err := e.errorTemplate.ExecuteTemplate(w, "layout", map[string]interface{}{"Error": http.StatusText(status), "Status": status}); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}
