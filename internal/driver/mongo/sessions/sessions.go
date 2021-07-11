package sessions

import (
	"context"
	"fmt"
	"net/http"
	"time"

	e "github.com/VxVxN/mdserver/pkg/error"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

const timeout = 5

type MongoSessions struct {
	mongoClient *mongo.Client
}

type Session struct {
	Token  string    `bson:"token"`
	Create time.Time `bson:"create"`
	Update time.Time `bson:"update"`
}

func Init(mongoClient *mongo.Client) *MongoSessions {
	return &MongoSessions{mongoClient: mongoClient}
}

func (ms *MongoSessions) getCollection() *mongo.Collection {
	return ms.mongoClient.Database("mdServer").Collection("sessions")
}

func (ms *MongoSessions) Create(token string) *e.ErrObject {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	_, err := ms.getCollection().InsertOne(ctx, Session{
		Token:  token,
		Create: time.Now(),
		Update: time.Now(),
	})
	if err != nil {
		return e.NewError("Can't add session to mongo", http.StatusInternalServerError, err)
	}

	return nil
}

func (ms *MongoSessions) Get(token string) (*Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	result := ms.getCollection().FindOne(ctx, bson.D{{"token", token}})

	var session Session
	err := result.Decode(&session)
	if err != nil {
		err = fmt.Errorf("error decode session: %v", err)
		return nil, err
	}

	return &session, nil
}

func (ms *MongoSessions) Delete(token string) *e.ErrObject {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	filter := bson.D{{"token", token}}
	_, err := ms.getCollection().DeleteOne(ctx, filter)
	if err == mongo.ErrNoDocuments {
		return e.NewError("Session not found", http.StatusNotFound, err)
	} else if err != nil {
		return e.NewError("Can't delete Session from mongo", http.StatusInternalServerError, err)
	}

	return nil
}
