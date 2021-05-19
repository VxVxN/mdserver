package post

import (
	"html/template"
	"mdserver/internal/glob"
	"mdserver/internal/post"
	"path"
)

type Controller struct {
	postTemplate *template.Template
	posts        *post.PostArray
}

func NewController() *Controller {
	return &Controller{
		postTemplate: template.Must(template.ParseFiles(path.Join(glob.WorkDir, "templates", "layout.html"), path.Join(glob.WorkDir, "templates", "post.html"))),
		posts:        post.NewPostArray(),
	}
}
