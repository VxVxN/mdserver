package share

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/log"

	"github.com/VxVxN/mdserver/internal/post"
	"github.com/VxVxN/mdserver/pkg/consts"
	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/VxVxN/mdserver/pkg/tools"
)

/**
 * @api {post} /share/:username/:id Get a share link
 * @apiName GetSharePostHandler
 * @apiGroup share
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 * {
 *    "message":"Failed to get post"
 * }
 */

func (ctrl *Controller) GetSharePostHandler(c *gin.Context) {
	templatePost, errObj := ctrl.getTemplateSharePost(c)
	if errObj != nil {
		log.Error.Printf("Failed to share post: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	c.HTML(200, "share_post.tmpl", templatePost)
}

func (ctrl *Controller) getTemplateSharePost(c *gin.Context) (*post.TemplatePost, *e.ErrObject) {
	username := strings.ToLower(c.Param("username"))

	shareLinks, err := ctrl.mongoShare.GetLinks(username)
	if err != nil {
		return nil, e.NewError("Failed to get share link", http.StatusInternalServerError, err)
	}

	id := c.Param("id")
	var pathToFile string
	for _, link := range shareLinks.ShareLinks {
		if link.ID == id {
			pathToFile = path.Join(consts.PathToPosts, username, link.DirName, link.FileName+consts.ExtMd)
			break
		}
	}
	templatePost, status, err := ctrl.posts.Get(pathToFile, false)
	if err != nil {
		err = fmt.Errorf("can't get post: %v", err)
		return nil, e.NewError("Failed to get post", status, err)
	}
	templatePost.IsAuth = tools.IsAuthUser(c)

	return &templatePost, nil
}
