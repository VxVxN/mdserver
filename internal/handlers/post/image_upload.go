package post

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"

	"github.com/VxVxN/mdserver/pkg/tools"

	"github.com/VxVxN/mdserver/pkg/consts"
	e "github.com/VxVxN/mdserver/pkg/error"

	"github.com/VxVxN/log"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) ImageUploadHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		log.Error.Printf("Incorrect multipart form data: %v", err)
		e.NewError("Incorrect multipart form data", http.StatusBadRequest, err)
		return
	}
	if errObj := ctrl.copyRequestFileToTmpFile(form); errObj != nil {
		log.Error.Printf("Failed to add tmp image: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) copyRequestFileToTmpFile(form *multipart.Form) *e.ErrObject {
	for _, header := range form.File {
		for _, fileHeader := range header {
			errObject := ctrl.copyFile(fileHeader)
			if errObject != nil {
				return errObject
			}
		}
	}
	return nil
}

func (ctrl *Controller) copyFile(fileHeader *multipart.FileHeader) *e.ErrObject {
	pathToDir := path.Join(consts.PathToTmpImages, fileHeader.Filename)

	requestFile, err := fileHeader.Open()
	if err != nil {
		err = fmt.Errorf("can't get request file: %v", err)
		return e.NewError("Failed to get request file", http.StatusInternalServerError, err)
	}

	defer tools.Close(requestFile, "Failed to close file, when add tmp image")

	file, err := os.Create(pathToDir)
	if err != nil {
		err = fmt.Errorf("can't create tmp image: %v", err)
		return e.NewError("Failed to create tmp image", http.StatusInternalServerError, err)
	}
	defer tools.Close(file, "Failed to close file, when add tmp image")

	if _, err = io.Copy(file, requestFile); err != nil {
		err = fmt.Errorf("can't copy tmp image: %v", err)
		return e.NewError("Failed to copy tmp image", http.StatusInternalServerError, err)
	}
	return nil
}
