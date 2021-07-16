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

	body := ctrl.prepareHTML(c, dirs)

	c.HTML(200, "index.tmpl", post.TemplatePost{
		Title: "Notes",
		Body:  template.HTML(body),
	})

	return nil
}

// TODO: transfer this to front
func (ctrl *Controller) prepareHTML(c *gin.Context, dirs []*posts.Directory) string {
	var body string
	_, err := tools.CheckCookie(c, ctrl.mongoSessions)
	if err == nil {
		body = `<button type="button" class="btn btn-outline-dark align-self-start" href="#" data-bs-toggle="modal" data-bs-target="#createDirectoryModal">Create directory</button>`
		for _, dir := range dirs {
			createPostBtn := `<button type="button" class="btn btn-outline-dark createPost align-self-start mb-3" href="#" data-bs-toggle="modal" data-bs-target="#createPostModal" data-dirname="` + dir.DirName + `">Create file</button>`

			deleteDirBtn := `[<a href="#" class="deleteModal d-inline-block" data-type="directory" data-name="` + dir.DirName + `">Delete directory</a>]`
			deleteDirBtn = `<button type="button" class="btn btn-outline-danger deleteModal d-inline-block" data-type="directory" data-name="` + dir.DirName + `">
								<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash-fill" viewBox="0 0 16 16">
									<path d="M2.5 1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1H3v9a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4h.5a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1H2.5zm3 4a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 .5-.5zM8 5a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7A.5.5 0 0 1 8 5zm3 .5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 1 0z"></path>
								</svg>
								<span class="visually-hidden">Button</span>
                        	</button>`
			body += `<div class="d-inline-block"><h3 class="align-self-start d-inline-block me-2 my-3">` + dir.DirName + `</h3>` + deleteDirBtn + `</div>` + createPostBtn
			body += `<div>`
			for _, file := range dir.Files {
				dirNameWithoutSpace := strings.Replace(dir.DirName, " ", "+", -1)
				fileNameWithoutSpace := strings.Replace(file, " ", "+", -1)
				body += `<div class="mb-2 d-flex">`

				linkToPost := dirNameWithoutSpace + "/" + fileNameWithoutSpace
				body += `<a type="button" class="btn flex-grow-1 btn-outline-dark me-1" href="/` + linkToPost + `">` + file + `</a>`
				body += `<div class="btn-group" role="group" aria-label="Basic example">`
				body += `<a type="button" class="btn btn-outline-secondary" href="/edit/` + linkToPost + `">
                			<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-pencil-fill" viewBox="0 0 16 16">
  								<path d="M12.854.146a.5.5 0 0 0-.707 0L10.5 1.793 14.207 5.5l1.647-1.646a.5.5 0 0 0 0-.708l-3-3zm.646 6.061L9.793 2.5 3.293 9H3.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.207l6.5-6.5zm-7.468 7.468A.5.5 0 0 1 6 13.5V13h-.5a.5.5 0 0 1-.5-.5V12h-.5a.5.5 0 0 1-.5-.5V11h-.5a.5.5 0 0 1-.5-.5V10h-.5a.499.499 0 0 1-.175-.032l-.179.178a.5.5 0 0 0-.11.168l-2 5a.5.5 0 0 0 .65.65l5-2a.5.5 0 0 0 .168-.11l.178-.178z"></path>
							</svg>
                			<span class="visually-hidden">Button</span>
              			</a>`
				body += `<button type="button" class="btn btn-outline-danger deleteModal" data-type="file" data-name="` + file + `" data-dirname="` + dir.DirName + `">
                            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash-fill" viewBox="0 0 16 16">
                                <path d="M2.5 1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1H3v9a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4h.5a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1H2.5zm3 4a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 .5-.5zM8 5a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7A.5.5 0 0 1 8 5zm3 .5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 1 0z"></path>
                            </svg>
                            <span class="visually-hidden">Button</span>
                        </button>`
				body += `</div>`
				body += `</div>`
			}
			body += `</div>`
		}
	}

	return body
}
