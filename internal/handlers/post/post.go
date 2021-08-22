package post

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/mdserver/internal/post"
	"github.com/VxVxN/mdserver/pkg/consts"
	"github.com/VxVxN/mdserver/pkg/tools"

	"github.com/VxVxN/log"

	e "github.com/VxVxN/mdserver/pkg/error"
)

func (ctrl *Controller) PostHandler(c *gin.Context) {
	templatePost, errObj := ctrl.getPost(c)
	if errObj != nil {
		log.Error.Printf("Failed to get post: %v", errObj.Error)
		if errObj.Status == http.StatusNotFound {
			c.HTML(http.StatusNotFound, "error.tmpl", map[string]interface{}{
				"Status": http.StatusNotFound,
				"Error":  "Page not found",
			})
		}
		errObj.JsonResponse(c)
		return
	}

	c.HTML(200, "post.tmpl", templatePost)
}

func (ctrl *Controller) getPost(c *gin.Context) (*post.TemplatePost, *e.ErrObject) {
	postMD, err := ctrl.getPathToPostMD(c)
	if err != nil {
		err = fmt.Errorf("can't get username from session: %v", err)
		return nil, e.NewError("Failed to get username from session", http.StatusBadRequest, err)
	}

	templatePost, status, err := ctrl.posts.Get(postMD, false)
	if err != nil {
		err = fmt.Errorf("can't get post: %v", err)
		return nil, e.NewError("Failed to get post", status, err)
	}
	return &templatePost, nil
}

func (ctrl *Controller) getPathToPostMD(c *gin.Context) (string, error) {
	dir := c.Param("dir")
	dir = strings.Replace(dir, "+", " ", -1)

	file := c.Param("file")
	file = strings.Replace(file, "+", " ", -1)

	username, err := tools.GetUserNameFromSession(c)
	if err != nil {
		return "", fmt.Errorf("cannot get username from session: %v", err)
	}

	postMD := path.Join(consts.PathToPosts, username, dir, file)
	postMD += consts.ExtMd
	return postMD, nil
}
