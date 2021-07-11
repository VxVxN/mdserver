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

	if errObj = SavePost(req.DirName, req.FileName, req.Text, false); errObj != nil {
		log.Error.Printf("Failed to save post: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func SavePost(dirName, fileName, text string, isCreateFile bool) *e.ErrObject {
	dirName = strings.Replace(dirName, "+", " ", -1)
	fileName = strings.Replace(fileName, "+", " ", -1)

	pathToFile := path.Join(consts.PathToPosts, dirName, fileName) + consts.ExtMd

	flags := os.O_TRUNC | os.O_WRONLY
	if isCreateFile {
		flags |= os.O_CREATE
	}

	file, err := os.OpenFile(pathToFile, flags, 0644)
	if err != nil {
		return e.NewError("Can't open file", http.StatusBadRequest, err)
	}
	defer tools.CloseFile(file)

	_, err = file.WriteString(text)
	if err != nil {
		return e.NewError("Error writing the file", http.StatusInternalServerError, err)
	}
	return nil
}
