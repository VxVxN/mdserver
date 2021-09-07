package login

import (
	"net/http"
	"strings"

	"github.com/VxVxN/log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/VxVxN/mdserver/pkg/tools"
)

type RequestSignIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ctrl *Controller) SignIn(c *gin.Context) {
	var req RequestSignIn
	session := sessions.Default(c)

	errObj := tools.UnmarshalRequest(c, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	if _, err := ctrl.mongoUsers.Get(req.Username, req.Password); err != nil {
		e.NewError("Incorrect login or password", http.StatusBadRequest, nil).JsonResponse(c)
		return
	}
	sessionToken := uuid.NewV4().String()

	session.Set("token", sessionToken)
	session.Set("username", strings.ToLower(req.Username))

	if err := session.Save(); err != nil {
		e.NewError("Failed to save session", http.StatusInternalServerError, nil).JsonResponse(c)
		return
	}
}
