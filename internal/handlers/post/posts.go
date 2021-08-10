package post

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/mdserver/pkg/tools"

	"github.com/VxVxN/mdserver/internal/post"

	"github.com/VxVxN/log"

	e "github.com/VxVxN/mdserver/pkg/error"
)

func (ctrl *Controller) PostsHandler(c *gin.Context) {
	if errObj := ctrl.getPosts(c); errObj != nil {
		log.Error.Printf("Failed to get edit post: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}
}

func (ctrl *Controller) getPosts(c *gin.Context) *e.ErrObject {
	directories, err := ctrl.prepareHTML(c)
	if err != nil {
		return e.NewError("Failed to prepare html", http.StatusBadRequest, err)
	}

	c.HTML(200, "index.tmpl", post.TemplatePosts{
		Title:       "Notes",
		IsAuth:      tools.IsAuthUser(c),
		Directories: directories,
	})

	return nil
}

func (ctrl *Controller) prepareHTML(c *gin.Context) ([]post.Directory, error) {
	var directories []post.Directory
	username, err := tools.GetUserNameFromSession(c)
	if err == nil {
		dirs, err := ctrl.mongoPosts.GetDirectories(username)
		if err != nil {
			return nil, fmt.Errorf("can't get posts for user: %s, error: %v", username, err)
		}
		for _, dir := range dirs {
			files := make([]post.File, 0, len(dir.Files))
			for _, file := range dir.Files {
				dirNameWithoutSpace := strings.Replace(dir.DirName, " ", "+", -1)
				fileNameWithoutSpace := strings.Replace(file.Name, " ", "+", -1)
				linkToPost := dirNameWithoutSpace + "/" + fileNameWithoutSpace
				files = append(files, post.File{
					Name:       file.Name,
					LinkToPost: linkToPost,
				})
			}
			directories = append(directories, post.Directory{
				DirName: dir.DirName,
				Files:   files,
			})
		}
	}
	return directories, nil
}
