package users

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/VxVxN/mdserver/internal/driver/mongo/interfaces/client"
	e "github.com/VxVxN/mdserver/pkg/error"
)

const timeout = 5

type MongoUsers struct {
	db client.Database
}

type User struct {
	Username string    `bson:"username"`
	Password string    `bson:"password"`
	Create   time.Time `bson:"create"`
}

func Init(db client.Database) *MongoUsers {
	return &MongoUsers{db: db}
}

func (ms *MongoUsers) getCollection() client.Collection {
	return ms.db.Collection("users")
}

func (ms *MongoUsers) Create(username, password string) *e.ErrObject {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	username = strings.ToLower(username)

	_, err := ms.getCollection().InsertOne(ctx, User{
		Username: username,
		Password: password,
		Create:   time.Now(),
	})
	if err != nil {
		return e.NewError("Can't add user to mongo", http.StatusInternalServerError, err)
	}

	return nil
}

func (ms *MongoUsers) Get(username, password string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	username = strings.ToLower(username)

	result := ms.getCollection().FindOne(ctx, bson.D{{"username", username}, {"password", password}})

	var user User
	if err := result.Decode(&user); err != nil {
		err = fmt.Errorf("error decode user: %v", err)
		return nil, err
	}

	return &user, nil
}
