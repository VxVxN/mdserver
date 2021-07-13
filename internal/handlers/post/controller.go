package post

import (
	"github.com/VxVxN/mdserver/internal/driver/mongo/sessions"

	"github.com/VxVxN/mdserver/internal/driver/mongo/posts"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/VxVxN/mdserver/internal/post"
	e "github.com/VxVxN/mdserver/pkg/error"
)

type Controller struct {
	e.ErrResponseController

	mongoPosts    *posts.MongoPosts
	mongoSessions *sessions.MongoSessions

	posts *post.Array
}

func NewController(mongoClient *mongo.Client) *Controller {
	ctrl := &Controller{
		mongoPosts:    posts.Init(mongoClient),
		mongoSessions: sessions.Init(mongoClient),
		posts:         post.NewPostArray(),
	}
	return ctrl
}
