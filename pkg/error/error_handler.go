package error

import (
	"html/template"
	"log"
	"mdserver/internal/glob"
	"net/http"
	"path"
)

var errorTemplate *template.Template

func init() {
	errorTemplate = template.Must(template.ParseFiles(path.Join(glob.WorkDir, "templates", "layout.html"), path.Join(glob.WorkDir, "templates", "error.html")))

}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	log.Printf("error %d %s %s\n", status, r.RemoteAddr, r.URL.Path)
	w.WriteHeader(status)
	if err := errorTemplate.ExecuteTemplate(w, "layout", map[string]interface{}{"Error": http.StatusText(status), "Status": status}); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
}
