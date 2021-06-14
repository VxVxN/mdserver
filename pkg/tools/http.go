package tools

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/log"

	e "github.com/VxVxN/mdserver/pkg/error"
)

func UnmarshalRequest(c *gin.Context, reqStruct interface{}) *e.ErrObject {
	if err := c.BindJSON(&reqStruct); err != nil {
		return e.NewError("Bad Request", http.StatusBadRequest, err)
	}

	return nil
}

func SuccessResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Error.Printf("Failed to unmarshal request: %v", err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Error.Printf("Failed to write json response: %v", err)
	}
}
