package tools

import (
	"net/http"

	"github.com/VxVxN/mdserver/internal/driver/mongo/sessions"

	"github.com/gin-gonic/gin"

	e "github.com/VxVxN/mdserver/pkg/error"
)

func UnmarshalRequest(c *gin.Context, reqStruct interface{}) *e.ErrObject {
	if err := c.BindJSON(&reqStruct); err != nil {
		return e.NewError("Bad Request", http.StatusBadRequest, err)
	}

	return nil
}

func CheckCookie(c *gin.Context, mongoSessions *sessions.MongoSessions) (int, error) {
	token, err := c.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return http.StatusUnauthorized, err
		}
		return http.StatusBadRequest, err
	}

	_, err = mongoSessions.Get(token)
	if err != nil {
		return http.StatusUnauthorized, err
	}

	return 0, nil
}
