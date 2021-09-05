package post

import (
	"fmt"
	"regexp"

	"github.com/VxVxN/mdserver/internal/driver/mongo/posts"
	"github.com/VxVxN/mdserver/internal/post"
	e "github.com/VxVxN/mdserver/pkg/error"
)

type Controller struct {
	e.ErrResponseController

	mongoPosts *posts.MongoPosts

	posts *post.Array

	imageLinkRegexp *regexp.Regexp
}

func NewController(mongoPosts *posts.MongoPosts) (*Controller, error) {
	imageLinkRegexp, err := regexp.Compile("!\\[]\\(\\/images\\/.*-.*-.*-.*\\)")
	if err != nil {
		return nil, fmt.Errorf("can't compile regexp: %v", err)
	}

	return &Controller{
		mongoPosts:      mongoPosts,
		posts:           post.NewPostArray(),
		imageLinkRegexp: imageLinkRegexp,
	}, nil
}
