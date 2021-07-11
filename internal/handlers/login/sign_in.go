package login

import (
	"net/http"

	"github.com/VxVxN/log"
	"github.com/VxVxN/mdserver/pkg/tools"

	"github.com/VxVxN/mdserver/internal/glob"

	"github.com/VxVxN/mdserver/pkg/config"

	e "github.com/VxVxN/mdserver/pkg/error"

	"github.com/gin-gonic/gin"

	uuid "github.com/satori/go.uuid"
)

type Credentials struct {
	//Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ctrl *Controller) SignIn(c *gin.Context) {
	var req Credentials

	errObj := tools.UnmarshalRequest(c, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	if config.Cfg.Password != req.Password {
		e.NewError("Incorrect login or password", http.StatusUnauthorized, nil).JsonResponse(c)
		return
	}

	sessionToken := uuid.NewV4().String()

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("session_token", sessionToken, 120, "", "", true, false)
	glob.SessionToken = sessionToken
}
