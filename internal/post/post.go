package post

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/russross/blackfriday"
)

type TemplatePost struct {
	Title   string
	Body    template.HTML
	ModTime int64
}

type TemplatePosts struct {
	Title       string
	ModTime     int64
	IsAuth      bool
	Directories []Directory
}

type Directory struct {
	DirName string
	Files   []File
}

type File struct {
	Name       string
	LinkToPost string
}

type Array struct {
	Items map[string]TemplatePost
	sync.RWMutex
}

func NewPostArray() *Array {
	p := Array{}
	p.Items = make(map[string]TemplatePost)
	return &p
}

// Get Loads the markdown file and converts it to HTML
// Returns an object of the TemplatePost type
// If the path does not exist or is a directory, then we return an error
func (p *Array) Get(md string, isEdit bool) (TemplatePost, int, error) {
	info, err := os.Stat(md)
	if err != nil {
		if os.IsNotExist(err) {
			// файл не существует
			return TemplatePost{}, 404, err
		}
		return TemplatePost{}, http.StatusInternalServerError, err
	}
	if info.IsDir() {
		return TemplatePost{}, http.StatusNotFound, fmt.Errorf("dir")
	}
	val, ok := p.Items[md+strconv.FormatBool(isEdit)]
	if !ok || (ok && val.ModTime != info.ModTime().UnixNano()) {
		p.RLock()
		defer p.RUnlock()
		fileText, _ := ioutil.ReadFile(md)
		lines := strings.Split(string(fileText), "\n")
		body := getBody(lines, isEdit)

		p.Items[md] = TemplatePost{"", template.HTML(body), info.ModTime().UnixNano()}
	}
	mdPost := p.Items[md]
	return mdPost, http.StatusOK, nil
}

func getBody(lines []string, isEdit bool) string {
	body := strings.Join(lines, "\n")

	if !isEdit {
		body = string(blackfriday.MarkdownCommon([]byte(body)))
	}
	return body
}
