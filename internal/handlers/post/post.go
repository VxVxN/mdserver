package post

import (
	"fmt"
	"path"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/mdserver/pkg/consts"

	"github.com/VxVxN/log"
	e "github.com/VxVxN/mdserver/pkg/error"
)

func (ctrl *Controller) PostHandler(c *gin.Context) {
	if errObj := ctrl.getPost(c); errObj != nil {
		log.Error.Printf("Failed to edit post: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) getPost(c *gin.Context) *e.ErrObject {
	postMD := ctrl.getPathToPostMD(c)

	templatePost, status, err := ctrl.posts.Get(postMD, false)
	if err != nil {
		err = fmt.Errorf("can't get post: %v, post: %s", err, postMD)
		return e.NewError("Failed to get post", status, err)
	}
	c.HTML(200, "post.tmpl", templatePost)

	return nil
}

func (ctrl *Controller) getPathToPostMD(c *gin.Context) string {
	dir := c.Param("dir")
	dir = strings.Replace(dir, "+", " ", -1)

	file := c.Param("file")
	file = strings.Replace(file, "+", " ", -1)

	postMD := path.Join(consts.PathToPosts, dir, file)
	postMD += consts.ExtMd
	return postMD
}
