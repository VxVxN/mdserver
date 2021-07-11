package login

import (
	"net/http"

	"github.com/VxVxN/log"
	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) LogOut(c *gin.Context) {

	token, err := c.Cookie("session_token")
	if err != nil {
		log.Error.Printf("Failed to get cookie: %v", err)
		e.NewError("session_token not found", http.StatusBadRequest, err)
		return
	}

	if errObj := ctrl.mongoSessions.Delete(token); errObj != nil {
		log.Error.Printf("Failed to delete session from mongo: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}
