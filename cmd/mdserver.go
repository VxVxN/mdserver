package main

import (
	"context"
	"fmt"
	slog "log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/VxVxN/mdserver/pkg/tools"

	e "github.com/VxVxN/mdserver/pkg/error"

	"github.com/VxVxN/mdserver/internal/handlers/login"

	"github.com/VxVxN/mdserver/pkg/consts"

	"github.com/VxVxN/log"
	"github.com/VxVxN/mdserver/internal/glob"
	"github.com/VxVxN/mdserver/internal/handlers/post"
	"github.com/VxVxN/mdserver/pkg/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mdServer struct {
	router      *gin.Engine
	MongoClient *mongo.Client

	postCtrl  *post.Controller
	loginCtrl *login.Controller
}

func main() {
	server, err := InitServer()
	if err != nil {
		slog.Fatal("Failed to init md server", err)
	}
	defer server.Stop()

	server.router.Static("/static/", glob.WorkDir+"/public/static")
	server.router.StaticFile("/favicon.ico", glob.WorkDir+"/public/static/images/favicon.ico")

	server.router.LoadHTMLGlob(consts.PathToTemplates + "/*")

	server.router.POST("/sign_in", server.loginCtrl.SignIn)

	authRouter := server.router.Group("")
	authRouter.Use(authMiddleware())
	{
		authRouter.POST("/delete_post", server.postCtrl.DeletePostHandler)
		authRouter.POST("/create_post", server.postCtrl.CreatePostHandler)
		authRouter.POST("/save_post", server.postCtrl.SavePostHandler)
		authRouter.POST("/preview", server.postCtrl.PreviewPostHandler)

		authRouter.POST("/create_directory", server.postCtrl.CreateDirectoryHandler)
		authRouter.POST("/delete_directory", server.postCtrl.DeleteDirectoryHandler)

		authRouter.GET("/edit/:dir/:file", server.postCtrl.EditPostHandler)
		authRouter.GET("/:dir/:file", server.postCtrl.PostHandler)
	}

	server.router.GET("/", server.postCtrl.PostsHandler)

	listen := config.Cfg.Listen
	log.Info.Printf("Listening %s", listen)
	if err = server.router.RunTLS(listen, "https-server.crt", "https-server.key"); err != nil {
		log.Fatal.Printf("Failed to run router: %v", err)
	}
}

func InitServer() (*mdServer, error) {
	server := mdServer{router: gin.Default()}

	pathConfig := path.Join(glob.WorkDir, "mdserver.yaml")
	err := config.InitConfig(pathConfig)
	if err != nil {
		return nil, fmt.Errorf("can't read config: %v, path: %s", err, pathConfig)
	}

	pathLogs := path.Join(glob.WorkDir, "logs/md_server.log")

	if err = log.Init(pathLogs, getLevelLog(config.Cfg.LevelLog), false); err != nil {
		return nil, fmt.Errorf("can't init log: %v, path: %s", err, pathLogs)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal.Printf("Failed to connect to mongo db: %v", err)
		return nil, fmt.Errorf("can't connect to mongo db: %v", err)
	}

	server.MongoClient = client

	server.postCtrl = post.NewController(server.MongoClient)
	server.loginCtrl = login.NewController()

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

func getLevelLog(lvlLog config.LVLLog) log.LevelLog {
	switch lvlLog {
	case config.DebugLog:
		return log.DebugLog
	case config.TraceLog:
		return log.TraceLog
	}
	return log.CommonLog
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		status, err := tools.CheckCookie(c)
		if err != nil {
			if status == http.StatusUnauthorized {
				c.HTML(http.StatusUnauthorized, "unauthorized.html", nil)
			} else {
				e.NewError("Bad Request", http.StatusBadRequest, err).JsonResponse(c)
			}
			c.Abort()
		}
	}
}
