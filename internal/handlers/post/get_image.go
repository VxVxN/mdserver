package post

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/VxVxN/log"
	"github.com/gin-gonic/gin"

	"github.com/VxVxN/mdserver/pkg/consts"
	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/VxVxN/mdserver/pkg/tools"
)

/**
 * @api {get} /images/:image Get image
 * @apiName GetImageHandler
 * @apiGroup post
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 * {
 *    "message":"Page not found"
 * }
 */

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
	postMD, err := getPathToImage(c)
	if err != nil {
		err = fmt.Errorf("can't get username from session: %v", err)
		return e.NewError("Failed to get username from session", http.StatusBadRequest, err)
	}

	if _, err = os.Stat(postMD); err != nil {
		postMD, err = getImageFromTMPImages(c)
		if err != nil {
			err = fmt.Errorf("can't get image from tmp images: %v", err)
			return e.NewError("Image not found", http.StatusNotFound, err)
		}
	}

	c.File(postMD)

	return nil
}

func getPathToImage(c *gin.Context) (string, error) {
	image := c.Param("image")
	image = strings.Replace(image, "+", " ", -1)

	postMD := path.Join(consts.PathToImages, image)

	return postMD, nil
}

func getImageFromTMPImages(c *gin.Context) (string, error) {
	image := c.Param("image")
	image = strings.Replace(image, "+", " ", -1)

	username, err := tools.GetUserNameFromSession(c)
	if err != nil {
		return "", fmt.Errorf("cannot get username from session: %v", err)
	}

	pathToTMP := path.Join(tools.GetPathToTMPImages(username))

	files, err := tools.GetFileNamesInDir(pathToTMP)
	if err != nil {
		return "", fmt.Errorf("can't get files names from tmp images: %v", err)
	}
	for _, file := range files {
		if strings.HasPrefix(file, image) {
			return path.Join(pathToTMP, file), nil
		}
	}
	return "", errors.New("image not found")
}
