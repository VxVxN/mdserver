package error

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/VxVxN/mdserver/internal/glob"
)

type ErrResponseController struct {
	errorTemplate *template.Template
}

func (e *ErrResponseController) InitErrResponseController() {
	e.errorTemplate = template.Must(template.ParseFiles(path.Join(glob.WorkDir, "templates", "layout.html"), path.Join(glob.WorkDir, "templates", "error.html")))
}

func (e *ErrResponseController) ErrorResponse(w http.ResponseWriter, r *http.Request, status int) {
	log.Printf("error %d %s %s\n", status, r.RemoteAddr, r.URL.Path)
	w.WriteHeader(status)
	if err := e.errorTemplate.ExecuteTemplate(w, "layout", map[string]interface{}{"Error": http.StatusText(status), "Status": status}); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func JsonErrorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
