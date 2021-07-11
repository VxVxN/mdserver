package post

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/VxVxN/mdserver/pkg/tools"
	"github.com/gin-gonic/gin"

	"github.com/VxVxN/mdserver/internal/driver/mongo/posts"

	"github.com/VxVxN/mdserver/internal/post"

	"github.com/VxVxN/log"
	e "github.com/VxVxN/mdserver/pkg/error"
)

func (ctrl *Controller) PostsHandler(c *gin.Context) {
	if errObj := ctrl.getPosts(c); errObj != nil {
		log.Error.Printf("Failed to edit post: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) getPosts(c *gin.Context) *e.ErrObject {
	dirs, err := ctrl.mongoPosts.GetList()
	if err != nil {
		err = fmt.Errorf("can't get posts: %v", err)
		return e.NewError("Failed to get posts", http.StatusInternalServerError, err)
	}

	body := prepareHTML(c, dirs)

	if err = ctrl.indexTemplate.ExecuteTemplate(c.Writer, "layout", post.TemplatePost{
		Title: "Notes",
		Body:  template.HTML(body),
	}); err != nil {
		err = fmt.Errorf("can't execute template: %v", err)
		return e.NewError("Failed to execute template", http.StatusInternalServerError, err)
	}
	return nil
}

// TODO: transfer this to front
func prepareHTML(c *gin.Context, dirs []*posts.Directory) string {
	var body string
	_, err := tools.CheckCookie(c)
	if err == nil {
		body = "[<a href=\"#\" data-bs-toggle=\"modal\" data-bs-target=\"#createDirectoryModal\">Create directory</a>]"
		for _, dir := range dirs {
			createPostBtn := "[<a href=\"#\" class=\"createPost\" data-bs-toggle=\"modal\" data-bs-target=\"#createPostModal\" data-dirname=\"" + dir.DirName + "\">Create file</a>]"
			deletePostBtn := "[<a href=\"#\" class=\"deleteModal\" data-type=\"directory\" data-name=\"" + dir.DirName + "\">Delete directory</a>]"
			body += "<h3>" + dir.DirName + "</h3>" + createPostBtn + deletePostBtn
			body += "<ul>"
			for _, file := range dir.Files {
				dirNameWithoutSpace := strings.Replace(dir.DirName, " ", "+", -1)
				fileNameWithoutSpace := strings.Replace(file, " ", "+", -1)

				linkToPost := dirNameWithoutSpace + "/" + fileNameWithoutSpace
				body += "<li>"
				body += "<a href=\"/" + linkToPost + "\">" + file + "</a>"
				body += " [<a href=\"/edit/" + linkToPost + "\">Edit</a>]"
				body += " [<a href=\"#/" + linkToPost + "\" class=\"deleteModal\" data-type=\"file\" data-name=\"" + file + "\" data-dirname=\"" + dir.DirName + "\">Delete</a>]"
				body += "</li>"
			}
			body += "</ul>"
		}
	}

	return body
}
