package post

import (
	"html/template"
	"path"

	"github.com/VxVxN/mdserver/pkg/consts"

	"github.com/VxVxN/mdserver/internal/post"
	e "github.com/VxVxN/mdserver/pkg/error"
)

type Controller struct {
	e.ErrResponseController

	postTemplate        *template.Template
	editingPostTemplate *template.Template
	posts               *post.Array
}

func NewController() *Controller {
	pathToLayout := path.Join(consts.PathToTemplates, "layout.html")
	ctrl := &Controller{
		postTemplate:        template.Must(template.ParseFiles(pathToLayout, path.Join(consts.PathToTemplates, "post.html"))),
		editingPostTemplate: template.Must(template.ParseFiles(pathToLayout, path.Join(consts.PathToTemplates, "editing_post.html"))),
		posts:               post.NewPostArray(),
	}
	ctrl.InitErrResponseController()
	return ctrl
}
