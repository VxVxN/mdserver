package tools

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"

	e "github.com/VxVxN/mdserver/pkg/error"
)

func UnmarshalRequest(c *gin.Context, reqStruct interface{}) *e.ErrObject {
	if err := c.BindJSON(&reqStruct); err != nil {
		return e.NewError("Bad Request", http.StatusBadRequest, err)
	}

	return nil
}

func CheckSession(c *gin.Context) (int, error) {
	session := sessions.Default(c)

	token := session.Get("token")
	if token == nil {
		return http.StatusUnauthorized, errors.New("token not found")
	}

	return 0, nil
}

func GetUserNameFromSession(c *gin.Context) (string, error) {
	session := sessions.Default(c)

	usernameI := session.Get("username")
	if usernameI == nil {
		return "", errors.New("username not found in session")
	}

	username, ok := usernameI.(string)
	if !ok {
		return "", fmt.Errorf("incorrect username in session, expected string, actual: %T", usernameI)
	}

	return username, nil
}
