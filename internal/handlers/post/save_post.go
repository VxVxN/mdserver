package post

import (
	"net/http"
	"os"
	"path"

	"github.com/VxVxN/mdserver/internal/glob"
	"github.com/VxVxN/mdserver/pkg/consts"
	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/VxVxN/mdserver/pkg/tools"
)

type RequestSave struct {
	FileName string `json:"name"`
	Text     string `json:"text"`
}

func (ctrl *Controller) SavePostHandler(w http.ResponseWriter, r *http.Request) {
	var req RequestSave

	errObj := tools.UnmarshalRequest(r, &req)
	if errObj != nil {
		errObj.JsonResponse(w)
		return
	}

	if errObj = SavePost(req.FileName, req.Text, false); errObj != nil {
		errObj.JsonResponse(w)
		return
	}
}

func SavePost(fileName, text string, isCreateFile bool) *e.ErrObject {
	pathToFile := path.Join(glob.WorkDir, "posts", fileName) + consts.ExtMd
	flags := os.O_TRUNC | os.O_WRONLY
	if isCreateFile {
		flags |= os.O_CREATE
	}

	file, err := os.OpenFile(pathToFile, flags, 0644)
	if err != nil {
		return e.NewError("Can't open file", http.StatusBadRequest, err)
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		return e.NewError("Error writing the file", http.StatusInternalServerError, err)
	}
	return nil
}
