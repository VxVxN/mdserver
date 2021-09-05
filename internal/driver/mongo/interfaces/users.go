package interfaces

import (
	"github.com/VxVxN/mdserver/internal/driver/mongo/users"
	e "github.com/VxVxN/mdserver/pkg/error"
)

type MongoUsers interface {
	Create(username, password string) *e.ErrObject
	Get(username, password string) (*users.User, error)
}
