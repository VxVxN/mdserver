package post

import (
	"html/template"
	"mdserver/internal/post"
	"path"
)

type Controller struct {
	postTemplate  *template.Template
	errorTemplate *template.Template
	posts         *post.PostArray

	workDir string // TODO: get it out of here
}

func NewController(workDir string) *Controller {
	return &Controller{
		workDir:       workDir,
		postTemplate:  template.Must(template.ParseFiles(path.Join(workDir, "templates", "layout.html"), path.Join(workDir, "templates", "post.html"))),
		errorTemplate: template.Must(template.ParseFiles(path.Join(workDir, "templates", "layout.html"), path.Join(workDir, "templates", "error.html"))),
		posts:         post.NewPostArray(),
	}
}
