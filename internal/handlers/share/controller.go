package share

import (
	"github.com/VxVxN/mdserver/internal/driver/mongo/interfaces"
	"github.com/VxVxN/mdserver/internal/post"
	e "github.com/VxVxN/mdserver/pkg/error"
)

type Controller struct {
	e.ErrResponseController

	posts *post.Array

	mongoShare interfaces.MongoShare
}

func NewController(mongoShare interfaces.MongoShare) *Controller {
	return &Controller{
		posts:      post.NewPostArray(),
		mongoShare: mongoShare,
	}
}
