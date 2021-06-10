package post

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/VxVxN/mdserver/internal/glob"
	"github.com/VxVxN/mdserver/pkg/consts"

	"github.com/VxVxN/log"
	e "github.com/VxVxN/mdserver/pkg/error"
)

func (ctrl *Controller) PostHandler(w http.ResponseWriter, r *http.Request) {
	if errObj := ctrl.getPost(w, r); errObj != nil {
		log.Error.Printf("Failed to edit post: %v", errObj.Error)
		errObj.JsonResponse(w)
		return
	}
}

func (ctrl *Controller) getPost(w http.ResponseWriter, r *http.Request) *e.ErrObject {
	postMD := ctrl.getPathToPostMD(r)

	templatePost, status, err := ctrl.posts.Get(postMD, false)
	if err != nil {
		err = fmt.Errorf("can't get post: %v, post: %s", err, postMD)
		return e.NewError("Failed to get post", status, err)
	}
	if err = ctrl.postTemplate.ExecuteTemplate(w, "layout", templatePost); err != nil {
		err = fmt.Errorf("can't execute template: %v, post: %s", err, templatePost.Title)
		return e.NewError("Failed to get post", http.StatusInternalServerError, err)
	}
	return nil
}

func (ctrl *Controller) getPathToPostMD(r *http.Request) string {
	params := r.URL.Query()

	dir := params.Get(":dir")
	dir = strings.Replace(dir, "+", " ", -1)

	file := params.Get(":file")
	file = strings.Replace(file, "+", " ", -1)

	postMD := path.Join(glob.WorkDir, "..", "posts", dir, file)
	postMD += consts.ExtMd
	return postMD
}
