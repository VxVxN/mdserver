package post

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/log"

	"github.com/VxVxN/mdserver/internal/driver/mongo/posts"
	"github.com/VxVxN/mdserver/pkg/consts"
	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/VxVxN/mdserver/pkg/tools"
)

type RequestSave struct {
	DirName  string `json:"dir_name" binding:"required"`
	FileName string `json:"file_name" binding:"required"`
	Text     string `json:"text" binding:"required"`
}

func (ctrl *Controller) SavePostHandler(c *gin.Context) {
	var req RequestSave

	errObj := tools.UnmarshalRequest(c, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	if errObj = ctrl.SavePost(c, req.DirName, req.FileName, req.Text, false); errObj != nil {
		log.Error.Printf("Failed to save post: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) SavePost(c *gin.Context, dirName, fileName, text string, isCreateFile bool) *e.ErrObject {
	dirName = strings.Replace(dirName, "+", " ", -1)
	fileName = strings.Replace(fileName, "+", " ", -1)

	username, err := tools.GetUserNameFromSession(c)
	if err != nil {
		err = fmt.Errorf("cannot get username from session: %v", err)
		return e.NewError("Failed to get username from session", http.StatusBadRequest, err)
	}

	pathToFile := path.Join(consts.PathToPosts, username, dirName, fileName) + consts.ExtMd

	flags := os.O_TRUNC | os.O_WRONLY
	if isCreateFile {
		flags |= os.O_CREATE
	}

	images, err := ctrl.saveImages(c, text)
	if err != nil {
		return e.NewError("Error save image", http.StatusInternalServerError, err)
	}

	if errObj := ctrl.mongoPosts.AddImagesToPost(username, dirName, fileName, images); errObj != nil {
		return errObj
	}

	file, err := os.OpenFile(pathToFile, flags, 0644)
	if err != nil {
		return e.NewError("Can't open file", http.StatusBadRequest, err)
	}
	defer tools.Close(file, "Failed to close file when save post")

	_, err = file.WriteString(text)
	if err != nil {
		return e.NewError("Error writing the file", http.StatusInternalServerError, err)
	}
	return nil
}

func (ctrl *Controller) saveImages(c *gin.Context, text string) ([]posts.Image, error) {
	username, err := tools.GetUserNameFromSession(c)
	if err != nil {
		return nil, fmt.Errorf("cannot get username from session: %v", err)
	}
	pathToTmpImages := path.Join(consts.PathToPosts, username, "tmp_images")

	tmpImageNames, err := tools.GetFileNamesInDir(pathToTmpImages)
	if err != nil {
		return nil, fmt.Errorf("cannot get image names from tmp images: %v", err)
	}

	rawImageLinks := ctrl.imageLinkRegexp.FindAllString(text, -1)
	imageNames := make([]posts.Image, 0, len(rawImageLinks))
	for _, rawStr := range rawImageLinks {
		rawStr = strings.TrimSpace(rawStr)
		rawSlice := strings.Split(rawStr, "/")
		if len(rawSlice) != 3 {
			return nil, fmt.Errorf("invalid path %v", rawStr)
		}
		rawStr = rawSlice[2]
		imageUUID := rawStr[:len(rawStr)-1]

		imageName, err := getImageName(tmpImageNames, imageUUID)
		if err != nil {
			continue // file already exists
		}
		imageNames = append(imageNames, posts.Image{
			UUID: imageUUID,
			Name: imageName,
		})

		pathToSource := path.Join(pathToTmpImages, imageUUID+"_"+imageName)
		pathToDestination := path.Join(consts.PathToPosts, username, "images", imageUUID)
		if err = tools.CopyFile(pathToSource, pathToDestination); err != nil {
			log.Debug.Printf("Failed to copy file: %v", err)
		}
	}

	if err = os.RemoveAll(pathToTmpImages); err != nil {
		log.Warning.Printf("Can't remove tmp images: %v", err)
	}
	_ = os.MkdirAll(pathToTmpImages, 0777)
	return imageNames, nil
}

func getImageName(fileNames []string, imageUUID string) (string, error) {
	for _, name := range fileNames {
		if strings.HasPrefix(name, imageUUID) {
			splitName := strings.Split(name, "_")
			if len(splitName) < 2 {
				return "", fmt.Errorf("invalid name: %s", name)
			}
			return strings.Join(splitName[1:], "_"), nil
		}
	}
	return "", fmt.Errorf("image not found")
}
