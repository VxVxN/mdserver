package post

import (
	"net/http"

	e "github.com/VxVxN/mdserver/pkg/error"

	"github.com/russross/blackfriday"

	"github.com/VxVxN/log"
	"github.com/VxVxN/mdserver/pkg/tools"
)

type RequestPreview struct {
	Text string `json:"text"`
}

func (ctrl *Controller) PreviewPostHandler(w http.ResponseWriter, r *http.Request) {
	var req RequestPreview

	errObj := tools.UnmarshalRequest(r, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(w)
		return
	}

	preview := blackfriday.MarkdownCommon([]byte(req.Text))

	_, err := w.Write(preview)
	if err != nil {
		log.Error.Printf("Failed to write response: %v", err)
		errObj = e.NewError("Failed to write response", http.StatusInternalServerError, err)
		errObj.JsonResponse(w)
		return
	}
}
