package error

import (
	"net/http"
)

type ErrResponseController struct {
}

func (e *ErrResponseController) ErrorResponse(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	//if err := e.errorTemplate.ExecuteTemplate(w, "layout", map[string]interfaces{}{"Error": http.StatusText(status), "Status": status}); err != nil {
	//	http.Error(w, http.StatusText(500), 500)
	//	return
	//}
}
