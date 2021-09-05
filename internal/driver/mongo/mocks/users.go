package mocks

import (
	"github.com/VxVxN/mdserver/internal/driver/mongo/users"
	e "github.com/VxVxN/mdserver/pkg/error"
)

type MongoUsers struct{}

func (ms *MongoUsers) Create(username, password string) *e.ErrObject {
	return nil
}

func (ms *MongoUsers) Get(username, password string) (*users.User, error) {
	return &users.User{
		Username: username,
		Password: password,
	}, nil
}
