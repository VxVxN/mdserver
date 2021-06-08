package main

import (
	slog "log"
	"net/http"
	"path"
	"strings"

	"github.com/VxVxN/mdserver/internal/handlers/common"

	"github.com/VxVxN/log"
	"github.com/VxVxN/mdserver/internal/glob"
	"github.com/VxVxN/mdserver/internal/handlers/post"
	"github.com/VxVxN/mdserver/pkg/config"

	"github.com/bmizerany/pat"
)

func init() {
	pathConfig := path.Join(glob.WorkDir, "mdserver.yaml")

	err := config.InitConfig(pathConfig)
	if err != nil {
		slog.Fatalf("Failed to read config: %v, name: %s", err, pathConfig)
	}

	pathLogs := path.Join(glob.WorkDir, "logs/md_server.log")
	if err = log.Init(pathLogs, getLevelLog(config.Cfg.LevelLog), false); err != nil {
		slog.Fatalf("Failed to init log: %v", err)
	}
}

func main() {
	// для отдачи сервером статичных файлов из папки public/static
	fs := noDirListing(http.FileServer(http.Dir(glob.WorkDir + "/public/static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	uploads := noDirListing(http.FileServer(http.Dir("./public/uploads")))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", uploads))

	postCtrl := post.NewController()
	commonCtrl := common.NewController()

	mux := pat.New()

	// static
	mux.Get("/favicon.ico", http.HandlerFunc(faviconHandler))

	// ajax
	mux.Post("/save", http.HandlerFunc(postCtrl.SavePostHandler))
	mux.Post("/save/", http.HandlerFunc(postCtrl.SavePostHandler))
	mux.Post("/preview", http.HandlerFunc(postCtrl.PreviewPostHandler))
	mux.Post("/preview/", http.HandlerFunc(postCtrl.PreviewPostHandler))

	mux.Post("/check_password", http.HandlerFunc(commonCtrl.CheckPasswordHandler))
	mux.Post("/check_password/", http.HandlerFunc(commonCtrl.CheckPasswordHandler))

	mux.Get("/edit/:page", http.HandlerFunc(postCtrl.EditPostHandler))
	mux.Get("/edit/:page/", http.HandlerFunc(postCtrl.EditPostHandler))
	mux.Get("/:page", http.HandlerFunc(postCtrl.PostHandler))
	mux.Get("/:page/", http.HandlerFunc(postCtrl.PostHandler))
	mux.Get("/", http.HandlerFunc(postCtrl.PostHandler))

	http.Handle("/", mux)

	listen := config.Cfg.Listen
	log.Info.Printf("Listening %s", listen)

	if err := http.ListenAndServe(listen, nil); err != nil {
		log.Fatal.Printf("Failed to listen and serve: %v, address: %s", err, listen)
	}
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
	http.ServeFile(w, r, glob.WorkDir+"/public/static/images/favicon.ico")
}
