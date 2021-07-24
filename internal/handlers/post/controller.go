package post

import (
	"fmt"
	"regexp"

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

	imageLinkRegexp *regexp.Regexp
}

func NewController(mongoClient *mongo.Client) (*Controller, error) {
	imageLinkRegexp, err := regexp.Compile("!\\[]\\(/static/images/.*\\)")
	if err != nil {
		return nil, fmt.Errorf("can't compile regexp: %v", err)
	}

	return &Controller{
		mongoPosts:      posts.Init(mongoClient),
		mongoSessions:   sessions.Init(mongoClient),
		posts:           post.NewPostArray(),
		imageLinkRegexp: imageLinkRegexp,
	}, nil
}
