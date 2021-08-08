package login

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/VxVxN/mdserver/internal/driver/mongo/users"
	e "github.com/VxVxN/mdserver/pkg/error"
)

type Controller struct {
	e.ErrResponseController

	mongoUsers *users.MongoUsers
}

func NewController(mongoClient *mongo.Client) *Controller {
	ctrl := &Controller{
		mongoUsers: users.Init(mongoClient),
	}
	return ctrl
}
