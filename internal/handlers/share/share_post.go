package share

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/VxVxN/log"

	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/VxVxN/mdserver/pkg/tools"
)

type RequestSharePost struct {
	DirName  string `json:"dir_name" binding:"required"`
	FileName string `json:"file_name" binding:"required"`
}

func (ctrl *Controller) SharePostHandler(c *gin.Context) {
	var req RequestSharePost

	errObj := tools.UnmarshalRequest(c, &req)
	if errObj != nil {
		log.Error.Printf("Failed to unmarshal request: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	link, errObj := ctrl.sharePost(c, req)
	if errObj != nil {
		log.Error.Printf("Failed to share post: %v", errObj.Error)
		errObj.JsonResponse(c)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"link": link,
	})
}

func (ctrl *Controller) sharePost(c *gin.Context, req RequestSharePost) (string, *e.ErrObject) {
	username, err := tools.GetUserNameFromSession(c)
	if err != nil {
		return "", e.NewError("Username not found in session", http.StatusBadRequest, err)
	}

	link, errObj := ctrl.mongoShare.GenerateLink(username, req.DirName, req.FileName)
	if errObj != nil {
		return "", errObj
	}
	return link, nil
}
