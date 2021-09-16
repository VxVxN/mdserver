package share

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/log"

	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/VxVxN/mdserver/pkg/tools"
)

type RequestDeleteShareLink struct {
	Link string `json:"link" binding:"required"`
}

/**
 * @api {post} /delete_share_link Delete a share link
 * @apiName DeleteShareLinkHandler
 * @apiGroup share
 *
 * @apiParamExample {json} Request example:
 * {
 *    "link":"https://vxvxn.ddns.net/share/vxvxn/57afcf14-0b2b-4a1e-be73-d16fbd5df032"
 * }
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 * {
 *    "message":"Share link not found"
 * }
 */

func (ctrl *Controller) DeleteShareLinkHandler(c *gin.Context) {
	var req RequestDeleteShareLink

	errObj := tools.UnmarshalRequest(c, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	errObj = ctrl.deleteShareLink(c, req)
	if errObj != nil {
		log.Error.Printf("Failed to delete share link: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) deleteShareLink(c *gin.Context, req RequestDeleteShareLink) *e.ErrObject {
	username, err := tools.GetUserNameFromSession(c)
	if err != nil {
		return e.NewError("Username not found in session", http.StatusBadRequest, err)
	}

	errObj := ctrl.mongoShare.DeleteLink(username, req.Link)
	if errObj != nil {
		return errObj
	}
	return nil
}
