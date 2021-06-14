package post

import (
	"net/http"

	"github.com/gin-gonic/gin"

	e "github.com/VxVxN/mdserver/pkg/error"

	"github.com/russross/blackfriday"

	"github.com/VxVxN/log"
	"github.com/VxVxN/mdserver/pkg/tools"
)

type RequestPreview struct {
	Text string `json:"text"`
}

func (ctrl *Controller) PreviewPostHandler(c *gin.Context) {
	var req RequestPreview

	errObj := tools.UnmarshalRequest(c, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	preview := blackfriday.MarkdownCommon([]byte(req.Text))

	_, err := c.Writer.Write(preview)
	if err != nil {
		log.Error.Printf("Failed to write response: %v", err)
		errObj = e.NewError("Failed to write response", http.StatusInternalServerError, err)
		errObj.JsonResponse(c)
		return
	}
}
