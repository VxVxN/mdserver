package login

import (
	"github.com/VxVxN/mdserver/internal/driver/mongo/interfaces"
	e "github.com/VxVxN/mdserver/pkg/error"
)

type Controller struct {
	e.ErrResponseController

	mongoUsers interfaces.MongoUsers
}

func NewController(mongoUsers interfaces.MongoUsers) *Controller {
	ctrl := &Controller{
		mongoUsers: mongoUsers,
	}
	return ctrl
}
