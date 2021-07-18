package posts

import (
	"context"
	"fmt"
	"net/http"
	"time"

	e "github.com/VxVxN/mdserver/pkg/error"

	"github.com/VxVxN/mdserver/pkg/tools"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const timeout = 5

type MongoPosts struct {
	mongoClient *mongo.Client
}

type Directory struct {
	DirName string   `bson:"dir"`
	Files   []string `bson:"files"`
}

func Init(mongoClient *mongo.Client) *MongoPosts {
	return &MongoPosts{mongoClient: mongoClient}
}

func (mgPost *MongoPosts) getCollection() *mongo.Collection {
	return mgPost.mongoClient.Database("mdServer").Collection("posts")
}

func (mgPost *MongoPosts) GetList() ([]*Directory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	cur, err := mgPost.getCollection().Find(ctx, bson.D{})
	if err != nil {
		err = fmt.Errorf("error find posts: %v", err)
		return nil, err
	}
	defer tools.CloseMongoCursor(cur, ctx)
	var posts []*Directory

	err = cur.All(ctx, &posts)
	if err != nil {
		err = fmt.Errorf("error get all posts: %v", err)
		return nil, err
	}

	return posts, nil
}

func (mgPost *MongoPosts) CreateDirectory(dirName string) *e.ErrObject {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	_, err := mgPost.getCollection().InsertOne(ctx, Directory{
		DirName: dirName,
		Files:   []string{},
	})
	if err == mongo.ErrNoDocuments {
		return e.NewError("Directory not found", http.StatusNotFound, err)
	} else if err != nil {
		return e.NewError("Can't create directory to mongo", http.StatusInternalServerError, err)
	}

	return nil
}

func (mgPost *MongoPosts) DeleteDirectory(dirName string) *e.ErrObject {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	filter := bson.D{{"dir", dirName}}
	_, err := mgPost.getCollection().DeleteOne(ctx, filter)
	if err == mongo.ErrNoDocuments {
		return e.NewError("Directory not found", http.StatusNotFound, err)
	} else if err != nil {
		return e.NewError("Can't delete directory to mongo", http.StatusInternalServerError, err)
	}

	return nil
}

func (mgPost *MongoPosts) RenameDirectory(oldDirName, newDirName string) *e.ErrObject {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	filter := bson.D{{"dir", oldDirName}}
	update := bson.D{{"$set", bson.D{{"dir", newDirName}}}}
	_, err := mgPost.getCollection().UpdateOne(ctx, filter, update)
	if err == mongo.ErrNoDocuments {
		return e.NewError("Directory not found", http.StatusNotFound, err)
	} else if err != nil {
		return e.NewError("Can't rename directory to mongo", http.StatusInternalServerError, err)
	}

	return nil
}

func (mgPost *MongoPosts) CreatePost(dirName, fileName string) *e.ErrObject {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	filter := bson.D{{"dir", dirName}}
	update := bson.D{{"$push", bson.D{{"files", fileName}}}}

	_, err := mgPost.getCollection().UpdateOne(ctx, filter, update)
	if err == mongo.ErrNoDocuments {
		return e.NewError("Directory not found", http.StatusNotFound, err)
	} else if err != nil {
		return e.NewError("Can't create post to mongo", http.StatusInternalServerError, err)
	}

	return nil
}

func (mgPost *MongoPosts) RenamePost(dirName, oldFileName, newFileName string) *e.ErrObject {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	filter := bson.D{{"dir", dirName}, {"files", oldFileName}}
	update := bson.D{{"$set", bson.D{{"files.$", newFileName}}}}

	_, err := mgPost.getCollection().UpdateOne(ctx, filter, update)
	if err == mongo.ErrNoDocuments {
		return e.NewError("Directory not found", http.StatusNotFound, err)
	} else if err != nil {
		return e.NewError("Can't create post to mongo", http.StatusInternalServerError, err)
	}

	return nil
}

func (mgPost *MongoPosts) DeletePost(dirName, fileName string) *e.ErrObject {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	filter := bson.D{{"dir", dirName}}
	update := bson.D{{"$pull", bson.D{{"files", fileName}}}}
	_, err := mgPost.getCollection().UpdateOne(ctx, filter, update)
	if err == mongo.ErrNoDocuments {
		return e.NewError("Post not found", http.StatusNotFound, err)
	} else if err != nil {
		return e.NewError("Can't delete post from mongo", http.StatusInternalServerError, err)
	}

	return nil
}
