package error

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path"

	"github.com/VxVxN/mdserver/pkg/consts"

	"github.com/VxVxN/log"
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

func JsonErrorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Error.Printf("Failed to unmarshal request: %v", err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Error.Printf("Failed to write json response: %v", err)
	}
}
