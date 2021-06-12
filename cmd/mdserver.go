package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/VxVxN/log"
	"github.com/VxVxN/mdserver/internal/glob"
	"github.com/VxVxN/mdserver/internal/handlers/common"
	"github.com/VxVxN/mdserver/internal/handlers/post"
	"github.com/VxVxN/mdserver/pkg/config"
	"github.com/bmizerany/pat"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mdServer struct {
	MongoClient *mongo.Client
}

func main() {
	server, err := InitServer()
	if err != nil {
		log.Fatal.Printf("Failed to init md server: %v", err)
	}
	defer server.Stop()

	fs := noDirListing(http.FileServer(http.Dir(glob.WorkDir + "/../public/static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	uploads := noDirListing(http.FileServer(http.Dir("./public/uploads")))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", uploads))

	postCtrl := post.NewController(server.MongoClient)
	commonCtrl := common.NewController()

	mux := pat.New()

	// static
	mux.Get("/favicon.ico", http.HandlerFunc(faviconHandler))

	// ajax
	Post(mux, "/delete_post", postCtrl.DeletePostHandler)
	Post(mux, "/create_post", postCtrl.CreatePostHandler)
	Post(mux, "/save_post", postCtrl.SavePostHandler)
	Post(mux, "/preview", postCtrl.PreviewPostHandler)

	Post(mux, "/create_directory", postCtrl.CreateDirectoryHandler)
	Post(mux, "/delete_directory", postCtrl.DeleteDirectoryHandler)

	Post(mux, "/check_password", commonCtrl.CheckPasswordHandler)

	Get(mux, "/edit/:dir/:file", postCtrl.EditPostHandler)
	Get(mux, "/:dir/:file", postCtrl.PostHandler)

	mux.Get("/", http.HandlerFunc(postCtrl.PostsHandler))

	http.Handle("/", mux)

	listen := config.Cfg.Listen
	log.Info.Printf("Listening %s", listen)

	if err = http.ListenAndServe(listen, nil); err != nil {
		log.Fatal.Printf("Failed to listen and serve: %v, address: %s", err, listen)
	}
}

func InitServer() (*mdServer, error) {
	var server mdServer

	pathConfig := path.Join(glob.WorkDir, "..", "mdserver.yaml")
	err := config.InitConfig(pathConfig)
	if err != nil {
		return nil, fmt.Errorf("can't read config: %v, path: %s", err, pathConfig)
	}

	pathLogs := path.Join(glob.WorkDir, "..", "logs/md_server.log")

	if err = log.Init(pathLogs, getLevelLog(config.Cfg.LevelLog), false); err != nil {
		return nil, fmt.Errorf("can't init log: %v, path: %s", err, pathLogs)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal.Printf("Failed to connect to mongo db: %v", err)
		return nil, fmt.Errorf("can't connect to mongo db: %v", err)
	}

	server.MongoClient = client

	return &server, nil
}

func (server *mdServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.MongoClient.Disconnect(ctx); err != nil {
		log.Fatal.Printf("Failed to disconnect to mongo db: %v", err)
		os.Exit(1)
	}
}

func Post(mux *pat.PatternServeMux, route string, handler func(http.ResponseWriter, *http.Request)) {
	mux.Post(route, http.HandlerFunc(handler))
	mux.Post(route+"/", http.HandlerFunc(handler))
}

func Get(mux *pat.PatternServeMux, route string, handler func(http.ResponseWriter, *http.Request)) {
	mux.Get(route, http.HandlerFunc(handler))
	mux.Get(route+"/", http.HandlerFunc(handler))
}

// обертка для http.FileServer, чтобы она не выдавала список файлов
// например, если открыть http://127.0.0.1:3000/static/,
// то будет видно список файлов внутри каталога.
// noDirListing - вернет 404 ошибку в этом случае.
func noDirListing(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") || r.URL.Path == "" {
			log.Error.Printf("The path not found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	}
}

func getLevelLog(lvlLog config.LVLLog) log.LevelLog {
	switch lvlLog {
	case config.DebugLog:
		return log.DebugLog
	case config.TraceLog:
		return log.TraceLog
	}
	return log.CommonLog
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, glob.WorkDir+"/../public/static/images/favicon.ico")
}
