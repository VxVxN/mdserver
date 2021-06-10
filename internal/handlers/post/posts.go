package post

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/VxVxN/mdserver/internal/driver/mongo/posts"

	"github.com/russross/blackfriday"

	"github.com/VxVxN/mdserver/internal/post"

	"github.com/VxVxN/log"
	e "github.com/VxVxN/mdserver/pkg/error"
)

func (ctrl *Controller) PostsHandler(w http.ResponseWriter, r *http.Request) {
	if errObj := ctrl.getPosts(w); errObj != nil {
		log.Error.Printf("Failed to edit post: %v", errObj.Error)
		errObj.JsonResponse(w)
		return
	}
}

func (ctrl *Controller) getPosts(w http.ResponseWriter) *e.ErrObject {
	dirs, err := ctrl.mongoPosts.GetList()
	if err != nil {
		err = fmt.Errorf("can't get posts: %v", err)
		return e.NewError("Failed to get posts", http.StatusInternalServerError, err)
	}

	body := prepareHTML(dirs)

	if err = ctrl.postTemplate.ExecuteTemplate(w, "layout", post.TemplatePost{
		Title: "Записки",
		Body:  template.HTML(body),
	}); err != nil {
		err = fmt.Errorf("can't execute template: %v", err)
		return e.NewError("Failed to execute template", http.StatusInternalServerError, err)
	}
	return nil
}

func prepareHTML(dirs []*posts.Dir) string {
	var body string
	for _, dir := range dirs {
		body += "<h3>" + dir.DirName + "</h3>"

		var mdPosts string
		for _, file := range dir.Files {
			dirNameWithoutSpace := strings.Replace(dir.DirName, " ", "+", -1)
			fileNameWithoutSpace := strings.Replace(file, " ", "+", -1)

			linkToPost := dirNameWithoutSpace + "/" + fileNameWithoutSpace

			mdPosts += "* [" + file + "](/" + linkToPost + ") [[Редактировать](/edit/" + linkToPost + ")]\n"
		}
		body += string(blackfriday.MarkdownCommon([]byte(mdPosts)))
	}
	return body
}
