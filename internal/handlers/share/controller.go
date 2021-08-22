package share

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/VxVxN/mdserver/internal/driver/mongo/share"
	"github.com/VxVxN/mdserver/internal/post"
	e "github.com/VxVxN/mdserver/pkg/error"
)

type Controller struct {
	e.ErrResponseController

	posts *post.Array

	mongoShare *share.MongoShare
}

func NewController(mongoClient *mongo.Client) *Controller {
	return &Controller{
		posts:      post.NewPostArray(),
		mongoShare: share.Init(mongoClient),
	}
}
