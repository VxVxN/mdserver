package main

import (
	"html/template"
	"log"
	"mdserver/internal/config"
	"mdserver/internal/post"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/bmizerany/pat"
)

var (
	workDir, configFileName string
	postTemplate   *template.Template
	errorTemplate  *template.Template
	posts         = post.NewPostArray()
)

func init() {
	workDir = path.Dir(os.Args[0])
	configFileName = path.Join(workDir, "mdserver.yaml")

	postTemplate = template.Must(template.ParseFiles(path.Join(workDir, "templates", "layout.html"), path.Join(workDir, "templates", "post.html")))
	errorTemplate = template.Must(template.ParseFiles(path.Join(workDir, "templates", "layout.html"), path.Join(workDir, "templates", "error.html")))
}

func main() {
	cfg, err := config.ReadConfig(configFileName)
	if err != nil {
		log.Fatalf("Failed to read config: %v, name: %s", err, configFileName)
	}
	// для отдачи сервером статичных файлов из папки public/static
	fs := noDirListing(http.FileServer(http.Dir(workDir+"/public/static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	uploads := noDirListing(http.FileServer(http.Dir("./public/uploads")))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", uploads))

	mux := pat.New()
	mux.Get("/:page", http.HandlerFunc(postHandler))
	mux.Get("/:page/", http.HandlerFunc(postHandler))
	mux.Get("/", http.HandlerFunc(postHandler))

	http.Handle("/", mux)
	log.Printf("Listening %s...", cfg.Listen)
	log.Fatalln(http.ListenAndServe(cfg.Listen, nil))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	page := params.Get(":page")
	postMD := path.Join(workDir, "posts", page)

	if page == "" {
		// если page пусто, то выдаем главную
		postMD += "/index.md"
	}

	post, status, err := posts.Get(postMD)
	if err != nil {
		errorHandler(w, r, status)
		return
	}
	if err := postTemplate.ExecuteTemplate(w, "layout", post); err != nil {
		log.Println(err.Error())
		errorHandler(w, r, 500)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	log.Printf("error %d %s %s\n", status, r.RemoteAddr, r.URL.Path)
	w.WriteHeader(status)
	if err := errorTemplate.ExecuteTemplate(w, "layout", map[string]interface{}{"Error": http.StatusText(status), "Status": status}); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
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
