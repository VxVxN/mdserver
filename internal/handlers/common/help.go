package common

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/VxVxN/log"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"

	"github.com/VxVxN/mdserver/pkg/consts"
	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/VxVxN/mdserver/pkg/tools"
)

type TemplateHelp struct {
	Title  string
	Body   template.HTML
	IsAuth bool
}

/**
 * @api {get} /help Return help page
 * @apiName Help
 * @apiGroup Common
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 500 Not Found
 * {
 *    "message":"Failed to open file"
 * }
 */

func (ctrl *Controller) Help(c *gin.Context) {
	helpTemplate, errObj := ctrl.getHelpTemplate(c)
	if errObj != nil {
		log.Error.Printf("Failed to get help template: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
	c.HTML(200, "help.tmpl", *helpTemplate)
}

func (ctrl *Controller) getHelpTemplate(c *gin.Context) (*TemplateHelp, *e.ErrObject) {
	pathToHelp := path.Join(consts.PathToMarkdown, "help.md")
	file, err := os.Open(pathToHelp)
	if err != nil {
		return nil, e.NewError("Failed to open file", http.StatusInternalServerError, err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, e.NewError("Failed to read file", http.StatusInternalServerError, err)
	}
	body := template.HTML(blackfriday.MarkdownCommon(data))
	return &TemplateHelp{
		IsAuth: tools.IsAuthUser(c),
		Body:   body,
	}, nil
}
