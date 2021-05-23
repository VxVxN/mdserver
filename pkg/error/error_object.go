package error

import (
	"errors"
	"net/http"
	"strings"
)

type ErrObject struct {
	Massage string
	Status  int
	Error   error
}

func NewError(message string, status int, err error) *ErrObject {
	if err == nil {
		err = errors.New(strings.ToLower(message))
	}
	return &ErrObject{
		Massage: message,
		Status:  status,
		Error:   err,
	}
}

func (errObj *ErrObject) JsonResponse(w http.ResponseWriter) {
	JsonErrorResponse(w, errObj.Massage, errObj.Status)
}
