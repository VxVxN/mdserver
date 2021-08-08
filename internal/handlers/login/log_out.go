package login

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	e "github.com/VxVxN/mdserver/pkg/error"
)

func (ctrl *Controller) LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		e.NewError("Failed to logout", http.StatusInternalServerError, err).JsonResponse(c)
		return
	}
}
