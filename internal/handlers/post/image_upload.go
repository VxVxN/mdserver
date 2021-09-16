package post

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"

	uuid "github.com/satori/go.uuid"

	"github.com/VxVxN/mdserver/pkg/tools"

	"github.com/VxVxN/mdserver/pkg/consts"
	e "github.com/VxVxN/mdserver/pkg/error"

	"github.com/VxVxN/log"

	"github.com/gin-gonic/gin"
)

/**
 * @api {post} /edit/:dir/:file/image_upload Upload the image to temporary files
 * @apiName ImageUploadHandler
 * @apiGroup post
 *
 * @apiSuccessExample {json} Success response example:
 *     HTTP/1.1 200 OK
 * {
 *    "image":"e8d4d5cd-d975-4d19-9551-3e29d674ecc3"
 * }
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 * {
 *    "message":"Image not found"
 * }
 */

func (ctrl *Controller) ImageUploadHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		log.Error.Printf("Incorrect multipart form data: %v", err)
		e.NewError("Incorrect multipart form data", http.StatusBadRequest, err)
		return
	}
	imageUUID, errObj := ctrl.copyRequestFileToTmpFile(c, form)
	if errObj != nil {
		log.Error.Printf("Failed to add tmp image: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"image": imageUUID,
	})
}

func (ctrl *Controller) copyRequestFileToTmpFile(c *gin.Context, form *multipart.Form) (string, *e.ErrObject) {
	for _, header := range form.File {
		for _, fileHeader := range header {
			imageUUID, errObject := ctrl.copyFile(c, fileHeader)
			if errObject != nil {
				return "", errObject
			}
			return imageUUID, nil // processing only one file
		}
	}
	return "", e.NewError("Image not found", http.StatusNotFound, nil)
}

func (ctrl *Controller) copyFile(c *gin.Context, fileHeader *multipart.FileHeader) (string, *e.ErrObject) {
	username, err := tools.GetUserNameFromSession(c)
	if err != nil {
		err = fmt.Errorf("cannot get username from session: %v", err)
		return "", e.NewError("Failed to get username from session", http.StatusBadRequest, err)
	}

	imageUUID := uuid.NewV4().String()

	pathToDir := path.Join(consts.PathToPosts, username, "tmp_images", imageUUID+"_"+fileHeader.Filename)

	requestFile, err := fileHeader.Open()
	if err != nil {
		err = fmt.Errorf("can't get request file: %v", err)
		return "", e.NewError("Failed to get request file", http.StatusInternalServerError, err)
	}

	defer tools.Close(requestFile, "Failed to close file, when add tmp image")

	file, err := os.Create(pathToDir)
	if err != nil {
		err = fmt.Errorf("can't create tmp image: %v", err)
		return "", e.NewError("Failed to create tmp image", http.StatusInternalServerError, err)
	}
	defer tools.Close(file, "Failed to close file, when add tmp image")

	if _, err = io.Copy(file, requestFile); err != nil {
		err = fmt.Errorf("can't copy tmp image: %v", err)
		return "", e.NewError("Failed to copy tmp image", http.StatusInternalServerError, err)
	}
	return imageUUID, nil
}
