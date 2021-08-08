package post

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/mdserver/pkg/consts"
	"github.com/VxVxN/mdserver/pkg/tools"

	"github.com/VxVxN/log"

	e "github.com/VxVxN/mdserver/pkg/error"
)

func (ctrl *Controller) GetImageHandler(c *gin.Context) {
	if errObj := ctrl.getImage(c); errObj != nil {
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
}

func (ctrl *Controller) getImage(c *gin.Context) *e.ErrObject {
	postMD, err := ctrl.getPathToImage(c)
	if err != nil {
		err = fmt.Errorf("can't get username from session: %v", err)
		return e.NewError("Failed to get username from session", http.StatusBadRequest, err)
	}

	c.File(postMD)

	return nil
}

func (ctrl *Controller) getPathToImage(c *gin.Context) (string, error) {
	image := c.Param("image")
	image = strings.Replace(image, "+", " ", -1)

	username, err := tools.GetUserNameFromSession(c)
	if err != nil {
		return "", fmt.Errorf("cannot get username from session: %v", err)
	}

	postMD := path.Join(consts.PathToPosts, username, "images", image)
	return postMD, nil
}
