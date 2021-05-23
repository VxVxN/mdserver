package post

import (
	"html/template"
	"mdserver/internal/glob"
	"mdserver/internal/post"
	e "mdserver/pkg/error"
	"path"
)

type Controller struct {
	e.ErrResponseController

	postTemplate *template.Template
	posts        *post.PostArray
}

func NewController() *Controller {
	ctrl := &Controller{
		postTemplate: template.Must(template.ParseFiles(path.Join(glob.WorkDir, "templates", "layout.html"), path.Join(glob.WorkDir, "templates", "post.html"))),
		posts:        post.NewPostArray(),
	}
	ctrl.InitErrResponseController()
	return ctrl
}
