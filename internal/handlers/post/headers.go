package post

import (
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/russross/blackfriday"

	"github.com/VxVxN/mdserver/pkg/config"

	"github.com/VxVxN/log"

	"github.com/VxVxN/mdserver/pkg/tools"
)

type RequestHeaders struct {
	Dir  string `json:"dir" binding:"required"`
	File string `json:"file" binding:"required"`
	Text string `json:"text" binding:"required"`
}

type ResponseHeaders struct {
	Dir  string `json:"dir"`
	File string `json:"file"`
	Text string `json:"text"`
}

/**
 * @api {post} /headers Get list headers
 * @apiName HeadersHandler
 * @apiGroup post
 *
 * @apiParamExample {json} Request example:
 * {
 *    "text":"# Header"
 * }
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 500 Internal Server Error
 * {
 *    "message":"Failed to write response"
 * }
 */

var headersRegexp = regexp.MustCompile(`<h[1-6]>.*?</h[1-6]>`)

func (ctrl *Controller) HeadersHandler(c *gin.Context) {
	var req RequestHeaders

	errObj := tools.UnmarshalRequest(c, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	preview := blackfriday.MarkdownCommon([]byte(req.Text))

	var headers []string

	headersRegexp.ReplaceAllStringFunc(string(preview), func(s string) string {
		url := config.Cfg.GetURL() + "/" + req.Dir + "/" + req.File
		switch {
		case strings.HasPrefix(s, "<h1>"):
			headers = append(headers, url+"/header1")
		case strings.HasPrefix(s, "<h2>"):
			headers = append(headers, url+"/header2")
		case strings.HasPrefix(s, "<h3>"):
			headers = append(headers, url+"/header3")
		case strings.HasPrefix(s, "<h4>"):
			headers = append(headers, url+"/header4")
		case strings.HasPrefix(s, "<h5>"):
			headers = append(headers, url+"/header5")
		case strings.HasPrefix(s, "<h6>"):
			headers = append(headers, url+"/header6")
		}
		return "111"
	})
	log.Error.Println(headers)

	// _, err := c.Writer.Write(preview)
	// if err != nil {
	// 	log.Error.Printf("Failed to write response: %v", err)
	// 	errObj = e.NewError("Failed to write response", http.StatusInternalServerError, err)
	// 	errObj.JsonResponse(c)
	// 	return
	// }
}
