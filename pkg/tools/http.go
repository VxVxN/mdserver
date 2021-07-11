package tools

import (
	"errors"
	"net/http"

	"github.com/VxVxN/mdserver/internal/glob"

	"github.com/gin-gonic/gin"

	e "github.com/VxVxN/mdserver/pkg/error"
)

func UnmarshalRequest(c *gin.Context, reqStruct interface{}) *e.ErrObject {
	if err := c.BindJSON(&reqStruct); err != nil {
		return e.NewError("Bad Request", http.StatusBadRequest, err)
	}

	return nil
}

func CheckCookie(c *gin.Context) (int, error) {
	cookie, err := c.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return http.StatusUnauthorized, err
		}
		return http.StatusBadRequest, err
	}

	if cookie != glob.SessionToken {
		return http.StatusUnauthorized, errors.New("unauthorized")
	}
	return 0, nil
}
