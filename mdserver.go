package main

import (
	"log"
	"mdserver/internal/glob"
	"mdserver/internal/handlers/post"
	"mdserver/pkg/config"
	"net/http"
	"path"
	"strings"

	"github.com/bmizerany/pat"
)

func main() {
	configFileName := path.Join(glob.WorkDir, "mdserver.yaml")

	cfg, err := config.ReadConfig(configFileName)
	if err != nil {
		log.Fatalf("Failed to read config: %v, name: %s", err, configFileName)
	}
	// для отдачи сервером статичных файлов из папки public/static
	fs := noDirListing(http.FileServer(http.Dir(glob.WorkDir + "/public/static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	uploads := noDirListing(http.FileServer(http.Dir("./public/uploads")))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", uploads))

	postCtrl := post.NewController()

	mux := pat.New()
	mux.Get("/edit/:page", http.HandlerFunc(postCtrl.EditPostHandler))
	mux.Get("/edit/:page/", http.HandlerFunc(postCtrl.EditPostHandler))
	mux.Get("/:page", http.HandlerFunc(postCtrl.PostHandler))
	mux.Get("/:page/", http.HandlerFunc(postCtrl.PostHandler))
	mux.Get("/", http.HandlerFunc(postCtrl.PostHandler))

	http.Handle("/", mux)
	log.Printf("Listening %s...", cfg.Listen)
	log.Fatalln(http.ListenAndServe(cfg.Listen, nil))
}

// обертка для http.FileServer, чтобы она не выдавала список файлов
// например, если открыть http://127.0.0.1:3000/static/,
// то будет видно список файлов внутри каталога.
// noDirListing - вернет 404 ошибку в этом случае.
func noDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") || r.URL.Path == "" {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}
