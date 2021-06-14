package error

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
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

func (errObj *ErrObject) JsonResponse(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = errObj.Massage

	c.JSON(errObj.Status, resp)
}
