package posts

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/VxVxN/mdserver/internal/driver/mongo/interfaces/client"
	e "github.com/VxVxN/mdserver/pkg/error"

	"github.com/VxVxN/mdserver/pkg/tools"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const timeout = 5

type MongoPosts struct {
	db client.Database
}

type Directory struct {
	Owner   string `bson:"owner"`
	DirName string `bson:"dir"`
	Files   []File `bson:"files"`
}

type File struct {
	Name   string  `bson:"name"`
	Images []Image `bson:"images"`
}

type Image struct {
	UUID string `bson:"uuid"`
	Name string `bson:"name"`
}

func Init(db client.Database) *MongoPosts {
	return &MongoPosts{db: db}
}

func (mgPost *MongoPosts) getCollection() client.Collection {
	return mgPost.db.Collection("posts")
}

func (mgPost *MongoPosts) GetDirectories(username string) ([]*Directory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	cur, err := mgPost.getCollection().Find(ctx, bson.D{{"owner", username}})
	if err != nil {
		err = fmt.Errorf("error find posts: %v", err)
		return nil, err
	}
	defer tools.CloseMongoCursor(cur, ctx)
	var posts []*Directory

	err = cur.All(ctx, &posts)
	if err != nil {
		err = fmt.Errorf("cannot get all posts: %v", err)
		return nil, err
	}

	return posts, nil
}

func (mgPost *MongoPosts) GetDirectory(username, dirName string) (*Directory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	var err error
	result := mgPost.getCollection().FindOne(ctx, bson.D{{"owner", username}, {"dir", dirName}})
	if result.Err() != nil {
		err = fmt.Errorf("error find posts: %v", result.Err())
		return nil, err
	}

	var dir Directory
	if err = result.Decode(&dir); err != nil {
		err = fmt.Errorf("cannot get dir: %v", err)
		return nil, err
	}

	return &dir, nil
}

func (mgPost *MongoPosts) CreateDirectory(username, dirName string) *e.ErrObject {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	_, err := mgPost.getCollection().InsertOne(ctx, Directory{
		Owner:   username,
		DirName: dirName,
		Files:   []File{},
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

func (mgPost *MongoPosts) CreatePost(username, dirName, fileName string) *e.ErrObject {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	filter := bson.D{{"owner", username}, {"dir", dirName}}
	update := bson.D{{"$push",
		bson.D{{"files",
			File{Name: fileName, Images: []Image{}},
		}}}}

	_, err := mgPost.getCollection().UpdateOne(ctx, filter, update)
	if err == mongo.ErrNoDocuments {
		return e.NewError("Directory not found", http.StatusNotFound, err)
	} else if err != nil {
		return e.NewError("Can't create post to mongo", http.StatusInternalServerError, err)
	}

	return nil
}

func (mgPost *MongoPosts) RenamePost(username, dirName, oldFileName, newFileName string) *e.ErrObject {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	filter := bson.D{
		{"owner", username},
		{"dir", dirName},
		{"files.name", oldFileName}}
	update := bson.D{{"$set", bson.D{{"files.$.name", newFileName}}}}

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
	filter := bson.D{
		{"dir", dirName},
	}
	update := bson.D{{"$pull", bson.D{{"files", bson.D{{"name", fileName}}}}}}
	_, err := mgPost.getCollection().UpdateOne(ctx, filter, update)
	if err == mongo.ErrNoDocuments {
		return e.NewError("Post not found", http.StatusNotFound, err)
	} else if err != nil {
		return e.NewError("Can't delete post from mongo", http.StatusInternalServerError, err)
	}

	return nil
}

func (mgPost *MongoPosts) AddImagesToPost(username, dirName, fileName string, images []Image) *e.ErrObject {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	filter := bson.D{
		{"owner", username},
		{"dir", dirName},
		{"files.name", fileName}}
	update := bson.D{{"$addToSet",
		bson.D{{"files.$.images",
			bson.D{{"$each", images}},
		}}}}

	_, err := mgPost.getCollection().UpdateOne(ctx, filter, update)
	if err == mongo.ErrNoDocuments {
		return e.NewError("File not found", http.StatusNotFound, err)
	} else if err != nil {
		return e.NewError("Can't add images to post to mongo", http.StatusInternalServerError, err)
	}

	return nil
}

func (mgPost *MongoPosts) GetImages(username, dirName, fileName string) ([]Image, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	result := mgPost.getCollection().FindOne(ctx, bson.D{
		{"owner", username},
		{"dir", dirName},
		{"files.name", fileName},
	})
	if result.Err() != nil {
		err := fmt.Errorf("error find posts: %v", result.Err())
		return nil, err
	}
	var posts Directory
	err := result.Decode(&posts)
	if err != nil {
		err = fmt.Errorf("cannot get all posts: %v", err)
		return nil, err
	}

	for _, file := range posts.Files {
		if file.Name == fileName {
			return file.Images, nil
		}
	}

	return nil, fmt.Errorf("post not found")
}
