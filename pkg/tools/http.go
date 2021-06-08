package tools

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/VxVxN/log"

	e "github.com/VxVxN/mdserver/pkg/error"
)

func UnmarshalRequest(r *http.Request, reqStruct interface{}) *e.ErrObject {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		return e.NewError("Content Type is not application/json", http.StatusUnsupportedMediaType, nil)
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&reqStruct); err != nil {
		var unmarshalErr *json.UnmarshalTypeError
		if errors.As(err, &unmarshalErr) {
			return e.NewError("Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest, err)
		} else {
			return e.NewError("Bad Request "+err.Error(), http.StatusBadRequest, err)
		}
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
