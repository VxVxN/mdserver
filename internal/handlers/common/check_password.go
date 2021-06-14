package common

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/mdserver/pkg/config"

	"github.com/VxVxN/log"
	"github.com/VxVxN/mdserver/pkg/tools"
)

type RequestCheckPassword struct {
	Password string `json:"password"`
}

type ResponseCheckPassword struct {
	IsValid bool `json:"valid"`
}

func (ctrl *Controller) CheckPasswordHandler(c *gin.Context) {
	var req RequestCheckPassword

	errObj := tools.UnmarshalRequest(c, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	var isPasswordCorrect bool
	if config.Cfg.Password == req.Password {
		isPasswordCorrect = true
	}

	c.JSON(http.StatusOK, ResponseCheckPassword{IsValid: isPasswordCorrect})
}
