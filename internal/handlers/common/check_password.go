package common

import (
	"net/http"

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

func (ctrl *Controller) CheckPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var req RequestCheckPassword

	errObj := tools.UnmarshalRequest(r, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(w)
		return
	}

	var isPasswordCorrect bool
	if config.Cfg.Password == req.Password {
		isPasswordCorrect = true
	}

	tools.SuccessResponse(w, ResponseCheckPassword{IsValid: isPasswordCorrect})
}
