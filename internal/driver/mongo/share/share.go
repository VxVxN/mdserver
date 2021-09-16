package share

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/VxVxN/mdserver/internal/driver/mongo/interfaces/client"
	"github.com/VxVxN/mdserver/pkg/config"
	e "github.com/VxVxN/mdserver/pkg/error"
)

const timeout = 5

type MongoShare struct {
	db client.Database
}

type Share struct {
	Owner      string `bson:"owner"`
	ShareLinks []Link `bson:"share_links"`
}

type Link struct {
	ID       string    `bson:"id"`
	DirName  string    `bson:"dir"`
	FileName string    `bson:"file"`
	Create   time.Time `bson:"create"`
	// Expire   time.Time `bson:"expire"`
}

func Init(db client.Database) *MongoShare {
	return &MongoShare{db: db}
}

func (ms *MongoShare) getCollection() client.Collection {
	return ms.db.Collection("share")
}

func (ms *MongoShare) GenerateLink(username, dirName, fileName string) (string, *e.ErrObject) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	username = strings.ToLower(username)

	shareUUID := uuid.NewV4()

	filter := bson.M{"owner": username}
	update := bson.M{"$push": bson.M{"share_links": Link{
		ID:       shareUUID.String(),
		DirName:  dirName,
		FileName: fileName,
		Create:   time.Now(),
	}}}

	isUpsert := true
	_, err := ms.getCollection().UpdateOne(ctx, filter, update, &options.UpdateOptions{
		Upsert: &isUpsert,
	})
	if err == mongo.ErrNoDocuments {
		return "", e.NewError("Share link not found", http.StatusNotFound, err)
	} else if err != nil {
		return "", e.NewError("Can't create share link to mongo", http.StatusInternalServerError, err)
	}

	link := CreateShareLink(username, shareUUID.String())

	return link, nil
}

func (ms *MongoShare) GetLinks(username string) (*Share, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	username = strings.ToLower(username)

	result := ms.getCollection().FindOne(ctx, bson.M{"owner": username})

	var share Share
	if err := result.Decode(&share); err != nil {
		return new(Share), nil // share link not found in mongo
	}

	return &share, nil
}

func (ms *MongoShare) DeleteLink(username, link string) *e.ErrObject {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	username = strings.ToLower(username)

	id, err := getIDFromShareLink(link)
	if err != nil {
		return e.NewError("Can't get id from share link", http.StatusBadRequest, err)
	}

	filter := bson.M{"owner": username}
	update := bson.M{"$pull": bson.M{"share_links": bson.M{"id": id}}}

	_, err = ms.getCollection().UpdateOne(ctx, filter, update)
	if err == mongo.ErrNoDocuments {
		return e.NewError("Share link not found", http.StatusNotFound, err)
	} else if err != nil {
		return e.NewError("Cannot delete share link", http.StatusInternalServerError, err)
	}

	return nil
}

func CreateShareLink(username, id string) string {
	return "https://" + config.Cfg.Domain + "/share/" + username + "/" + id
}

func getIDFromShareLink(link string) (string, error) {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	splitPath := strings.Split(parsedURL.Path, "/")

	if len(splitPath) != 4 {
		return "", fmt.Errorf("invalid link: %s", link)
	}

	return splitPath[3], nil
}
