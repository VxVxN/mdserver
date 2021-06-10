package posts

import (
	"context"
	"fmt"
	"time"

	"github.com/VxVxN/mdserver/pkg/tools"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const timeout = 5

type MongoPosts struct {
	mongoClient *mongo.Client
}

type Dir struct {
	DirName string   `bson:"dir"`
	Files   []string `bson:"files"`
}

func Init(mongoClient *mongo.Client) *MongoPosts {
	return &MongoPosts{mongoClient: mongoClient}
}

func (mgPost *MongoPosts) getCollection() *mongo.Collection {
	return mgPost.mongoClient.Database("mdServer").Collection("posts")
}

//func (mgPost *MongoPosts) Get(dirName, fileName string) (*Post, *e.ErrObject) {
//	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
//	defer cancel()
//
//	var post Post
//
//	query := bson.D{{"dirname", dirName}, {"filename", fileName}}
//	err := mgPost.getCollection().FindOne(ctx, query).Decode(&post)
//	if err == mongo.ErrNoDocuments {
//		return nil, e.NewError("Post not found", http.StatusNotFound, err)
//	} else if err != nil {
//		return nil, e.NewError("Can't get post", http.StatusInternalServerError, err)
//	}
//
//	return &post, nil
//}

func (mgPost *MongoPosts) GetList() ([]*Dir, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	cur, err := mgPost.getCollection().Find(ctx, bson.D{})
	if err != nil {
		err = fmt.Errorf("error find posts: %v", err)
		return nil, err
	}
	defer tools.CloseMongoCursor(cur, ctx)
	var posts []*Dir

	err = cur.All(ctx, &posts)
	if err != nil {
		err = fmt.Errorf("error get all posts: %v", err)
		return nil, err
	}

	return posts, nil
}

//func (mgPost *MongoPosts) Insert() (*Post, *e.ErrObject) {
//	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
//	defer cancel()
//
//	var post Post
//
//	_, err := mgPost.getCollection().InsertOne(ctx, Post{})
//	if err == mongo.ErrNoDocuments {
//		return nil, e.NewError("Post not found", http.StatusNotFound, err)
//	} else if err != nil {
//		return nil, e.NewError("Can't get post", http.StatusInternalServerError, err)
//	}
//
//	return &post, nil
//}
