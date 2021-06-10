package post

import (
	"html/template"
	"path"

	"github.com/VxVxN/mdserver/internal/driver/mongo/posts"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/VxVxN/mdserver/pkg/consts"

	"github.com/VxVxN/mdserver/internal/post"
	e "github.com/VxVxN/mdserver/pkg/error"
)

type Controller struct {
	e.ErrResponseController

	mongoClient *mongo.Client
	mongoPosts  *posts.MongoPosts

	postTemplate        *template.Template
	editingPostTemplate *template.Template
	posts               *post.Array
}

func NewController(mongoClient *mongo.Client) *Controller {
	pathToLayout := path.Join(consts.PathToTemplates, "layout.html")
	ctrl := &Controller{
		mongoClient:         mongoClient,
		mongoPosts:          posts.Init(mongoClient),
		postTemplate:        template.Must(template.ParseFiles(pathToLayout, path.Join(consts.PathToTemplates, "post.html"))),
		editingPostTemplate: template.Must(template.ParseFiles(pathToLayout, path.Join(consts.PathToTemplates, "editing_post.html"))),
		posts:               post.NewPostArray(),
	}
	ctrl.InitErrResponseController()
	return ctrl
}
