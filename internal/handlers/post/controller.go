package post

import (
	"html/template"
	"path"

	"github.com/VxVxN/mdserver/internal/glob"
	"github.com/VxVxN/mdserver/internal/post"
	e "github.com/VxVxN/mdserver/pkg/error"
)

type Controller struct {
	e.ErrResponseController

	postTemplate        *template.Template
	editingPostTemplate *template.Template
	posts               *post.PostArray
}

func NewController() *Controller {
	ctrl := &Controller{
		postTemplate:        template.Must(template.ParseFiles(path.Join(glob.WorkDir, "templates", "layout.html"), path.Join(glob.WorkDir, "templates", "post.html"))),
		editingPostTemplate: template.Must(template.ParseFiles(path.Join(glob.WorkDir, "templates", "layout.html"), path.Join(glob.WorkDir, "templates", "editing_post.html"))),
		posts:               post.NewPostArray(),
	}
	ctrl.InitErrResponseController()
	return ctrl
}
