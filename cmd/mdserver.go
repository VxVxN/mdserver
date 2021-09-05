package main

import (
	"context"
	"fmt"
	slog "log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/VxVxN/log"
	ginSessions "github.com/gin-contrib/sessions"
	sessionMongo "github.com/gin-contrib/sessions/mongo"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"

	_interface "github.com/VxVxN/mdserver/internal/driver/mongo/interfaces/client"
	mongoClient "github.com/VxVxN/mdserver/internal/driver/mongo/mongo_client"
	"github.com/VxVxN/mdserver/internal/driver/mongo/posts"
	mShare "github.com/VxVxN/mdserver/internal/driver/mongo/share"
	"github.com/VxVxN/mdserver/internal/driver/mongo/users"
	"github.com/VxVxN/mdserver/internal/glob"
	"github.com/VxVxN/mdserver/internal/handlers/login"
	"github.com/VxVxN/mdserver/internal/handlers/post"
	"github.com/VxVxN/mdserver/internal/handlers/share"
	"github.com/VxVxN/mdserver/pkg/config"
	"github.com/VxVxN/mdserver/pkg/consts"
	e "github.com/VxVxN/mdserver/pkg/error"
	"github.com/VxVxN/mdserver/pkg/tools"
)

type mdServer struct {
	router      *gin.Engine
	MongoClient _interface.Client

	postCtrl  *post.Controller
	loginCtrl *login.Controller
	shareCtrl *share.Controller
}

func main() {
	server, err := InitServer()
	if err != nil {
		slog.Fatal("Failed to init md server", err)
	}
	defer server.Stop()

	server.router.Static("/static/", consts.PathToStatic)
	server.router.StaticFile("/favicon.ico", consts.PathToStaticImages+"/favicon.ico")

	server.router.LoadHTMLGlob(consts.PathToTemplates + "/*")

	server.router.POST("/sign_in", server.loginCtrl.SignIn)
	server.router.POST("/log_out", server.loginCtrl.LogOut)

	authRouter := server.router.Group("")
	authRouter.Use(server.authMiddleware())
	{
		authRouter.POST("/create_post", server.postCtrl.CreatePostHandler)
		authRouter.POST("/save_post", server.postCtrl.SavePostHandler)
		authRouter.POST("/rename_post", server.postCtrl.RenamePostHandler)
		authRouter.POST("/delete_post", server.postCtrl.DeletePostHandler)
		authRouter.POST("/preview", server.postCtrl.PreviewPostHandler)

		authRouter.POST("/share_link", server.shareCtrl.SharePostHandler)
		authRouter.POST("/delete_share_link", server.shareCtrl.DeleteShareLinkHandler)
		authRouter.GET("/share_links", server.shareCtrl.ShareLinksHandler)

		authRouter.POST("/create_directory", server.postCtrl.CreateDirectoryHandler)
		authRouter.POST("/rename_directory", server.postCtrl.RenameDirectoryHandler)
		authRouter.POST("/delete_directory", server.postCtrl.DeleteDirectoryHandler)

		authRouter.POST("/edit/:dir/:file/image_upload", server.postCtrl.ImageUploadHandler)
		authRouter.GET("/edit/:dir/:file", server.postCtrl.EditPostHandler)
		authRouter.GET("/:dir/:file", server.postCtrl.PostHandler)
	}

	server.router.GET("/", server.postCtrl.PostsHandler)
	server.router.GET("/share/:username/:id", server.shareCtrl.GetSharePostHandler)
	server.router.GET("/images/:image", server.postCtrl.GetImageHandler)

	server.router.NoRoute(noRouteHandler)

	listen := config.Cfg.Listen
	log.Info.Printf("Listening %s", listen)
	if err = server.router.RunTLS(listen, "https-server.crt", "https-server.key"); err != nil {
		log.Fatal.Printf("Failed to run router: %v", err)
	}
}

func InitServer() (*mdServer, error) {
	var err error
	gin.SetMode(gin.ReleaseMode)
	server := mdServer{router: gin.Default()}

	pathConfig := path.Join(glob.WorkDir, "mdserver.yaml")

	if err = config.InitConfig(pathConfig); err != nil {
		return nil, fmt.Errorf("can't read config: %v, path: %s", err, pathConfig)
	}

	pathLogs := path.Join(glob.WorkDir, "logs/md_server.log")

	if err = log.Init(pathLogs, getLevelLog(config.Cfg.LevelLog), false); err != nil {
		return nil, fmt.Errorf("can't init log: %v, path: %s", err, pathLogs)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := "mongodb://mongo:27017"

	client := mongoClient.NewClient()
	if err = client.Connect(ctx, mongoURI); err != nil {
		log.Fatal.Printf("Failed to connect to mongo db: %v", err)
		return nil, fmt.Errorf("can't connect to mongo db: %v", err)
	}
	server.MongoClient = client

	session, err := mgo.Dial(mongoURI)
	if err != nil {
		log.Fatal.Printf("Failed to connect to session mongo db: %v", err)
		return nil, fmt.Errorf("can't connect to session mongo db: %v", err)
	}

	c := session.DB("mdServer").C("sessions")
	store := sessionMongo.NewStore(c, config.Cfg.SessionAge*60, true, []byte(config.Cfg.SessionSecret))
	server.router.Use(ginSessions.Sessions("mdSession", store))

	mongoDB := server.MongoClient.Database("mdServer")
	mongoUsers := users.Init(mongoDB)
	mongoShare := mShare.Init(mongoDB)
	mongoPosts := posts.Init(mongoDB)

	server.postCtrl, err = post.NewController(mongoPosts)
	if err != nil {
		log.Fatal.Printf("Failed init post controller : %v", err)
		return nil, fmt.Errorf("error on init post controller: %v", err)
	}

	server.loginCtrl = login.NewController(mongoUsers)
	server.shareCtrl = share.NewController(mongoShare)

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

func (server *mdServer) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		status, err := tools.CheckSession(c)
		if err != nil {
			if status == http.StatusUnauthorized {
				c.HTML(http.StatusUnauthorized, "error.tmpl", map[string]interface{}{
					"Status": http.StatusUnauthorized,
					"Error":  "Unauthorized",
				})
			} else {
				e.NewError("Bad Request", http.StatusBadRequest, err).JsonResponse(c)
			}
			c.Abort()
		}
	}
}

func noRouteHandler(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error.tmpl", map[string]interface{}{
		"Status": http.StatusNotFound,
		"Error":  "Page not found",
	})
}
