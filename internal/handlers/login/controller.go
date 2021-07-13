package login

import (
	"github.com/VxVxN/mdserver/internal/driver/mongo/sessions"
	e "github.com/VxVxN/mdserver/pkg/error"
	"go.mongodb.org/mongo-driver/mongo"
)

type Controller struct {
	e.ErrResponseController

	mongoSessions *sessions.MongoSessions
}

func NewController(mongoClient *mongo.Client) *Controller {
	ctrl := &Controller{
		mongoSessions: sessions.Init(mongoClient),
	}
	return ctrl
}
