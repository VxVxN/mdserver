package post

import (
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/log"
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

	if errObj = ctrl.SavePost(req.DirName, req.FileName, req.Text, false); errObj != nil {
		log.Error.Printf("Failed to save post: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) SavePost(dirName, fileName, text string, isCreateFile bool) *e.ErrObject {
	dirName = strings.Replace(dirName, "+", " ", -1)
	fileName = strings.Replace(fileName, "+", " ", -1)

	pathToFile := path.Join(consts.PathToPosts, dirName, fileName) + consts.ExtMd

	flags := os.O_TRUNC | os.O_WRONLY
	if isCreateFile {
		flags |= os.O_CREATE
	}

	ctrl.saveImages(text, fileName)

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

func (ctrl *Controller) saveImages(text string, fileName string) {
	rawImageLinks := ctrl.imageLinkRegexp.FindAllString(text, -1)

	for _, rawStr := range rawImageLinks {
		rawStr = strings.TrimSpace(rawStr)
		rawStr = strings.TrimLeft(rawStr, "![](/static/images/")
		fileName = rawStr[:len(rawStr)-1]

		pathToSource := path.Join(consts.PathToTmpImages, fileName)
		pathToDestination := path.Join(consts.PathToStaticImages, fileName)

		_ = tools.Copy(pathToSource, pathToDestination)
	}

	if err := os.RemoveAll(consts.PathToTmpImages); err != nil {
		log.Warning.Printf("Can't remove tmp images: %v", err)
	}
	_ = os.MkdirAll(consts.PathToTmpImages, 0777)
}
