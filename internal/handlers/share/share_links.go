package share

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/log"

	"github.com/VxVxN/mdserver/internal/driver/mongo/share"
	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/VxVxN/mdserver/pkg/tools"
)

type ResponseShareLinks struct {
	Title      string
	ShareLinks []Link
}

type Link struct {
	Directory string
	File      string
	Link      string
}

/**
 * @api {get} /share_links Get list of sharing link
 * @apiName ShareLinksHandler
 * @apiGroup share
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 500 Internal Server Error
 * {
 *    "message":"Failed to get share link"
 * }
 */

func (ctrl *Controller) ShareLinksHandler(c *gin.Context) {

	resp, errObj := ctrl.getShareLinks(c)
	if errObj != nil {
		log.Error.Printf("Failed to get share links: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	c.HTML(200, "share_links.tmpl", resp)
}

func (ctrl *Controller) getShareLinks(c *gin.Context) (*ResponseShareLinks, *e.ErrObject) {
	username, err := tools.GetUserNameFromSession(c)
	if err != nil {
		return nil, e.NewError("Username not found in session", http.StatusBadRequest, err)
	}

	links, err := ctrl.mongoShare.GetLinks(username)
	if err != nil {
		return nil, e.NewError("Failed to get share link", http.StatusInternalServerError, err)
	}

	response := convertLinks(username, links)

	return response, nil
}

func convertLinks(username string, links *share.Share) *ResponseShareLinks {
	var resp ResponseShareLinks
	for _, link := range links.ShareLinks {
		resp.ShareLinks = append(resp.ShareLinks, Link{
			Directory: link.DirName,
			File:      link.FileName,
			Link:      share.CreateShareLink(username, link.ID),
		})
	}
	return &resp
}
