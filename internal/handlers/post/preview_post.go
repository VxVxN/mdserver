package post

import (
	"fmt"
	"net/http"
	"path"

	"github.com/VxVxN/mdserver/internal/glob"

	"github.com/VxVxN/mdserver/pkg/consts"

	"github.com/VxVxN/log"
	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/VxVxN/mdserver/pkg/tools"
)

type RequestPreview struct {
	FileName string `json:"name"`
}

func (ctrl *Controller) PreviewPostHandler(w http.ResponseWriter, r *http.Request) {
	var req RequestSave

	errObj := tools.UnmarshalRequest(r, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(w)
		return
	}

	if errObj = ctrl.previewPost(w, req.FileName); errObj != nil {
		log.Error.Printf("Failed to preview post: %v", errObj.Error)
		errObj.JsonResponse(w)
		return
	}
}

func (ctrl *Controller) previewPost(w http.ResponseWriter, fileName string) *e.ErrObject {
	fileName += consts.ExtMd
	postMD := path.Join(glob.WorkDir, "posts", fileName)

	templatePost, status, err := ctrl.posts.Get(postMD, false)
	if err != nil {
		err = fmt.Errorf("can't get post: %v, post: %s", err, fileName)
		return e.NewError("Failed to get preview post", status, err)
	}

	w.Write([]byte(templatePost.Body))

	return nil
}
