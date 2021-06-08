package post

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/russross/blackfriday"
)

type post struct {
	Title   string
	Body    template.HTML
	ModTime int64
}

type PostArray struct {
	Items map[string]post
	sync.RWMutex
}

func NewPostArray() *PostArray {
	p := PostArray{}
	p.Items = make(map[string]post)
	return &p
}

// Get Загружает markdown-файл и конвертирует его в HTML
// Возвращает объект типа Post
// Если путь не существует или является каталогом, то возвращаем ошибку
func (p *PostArray) Get(md string, isEdit bool) (post, int, error) {
	info, err := os.Stat(md)
	if err != nil {
		if os.IsNotExist(err) {
			// файл не существует
			return post{}, 404, err
		}
		return post{}, 500, err
	}
	if info.IsDir() {
		return post{}, 404, fmt.Errorf("dir")
	}
	val, ok := p.Items[md+strconv.FormatBool(isEdit)]
	if !ok || (ok && val.ModTime != info.ModTime().UnixNano()) {
		p.RLock()
		defer p.RUnlock()
		fileText, _ := ioutil.ReadFile(md)
		lines := strings.Split(string(fileText), "\n")
		body := getBody(lines, isEdit)
		var title string

		if md == "posts/index.md" {
			title = "Записки"
		}

		p.Items[md] = post{title, template.HTML(body), info.ModTime().UnixNano()}
	}
	mdPost := p.Items[md]
	return mdPost, 200, nil
}

func getBody(lines []string, isEdit bool) string {
	body := strings.Join(lines, "\n")

	if !isEdit {
		body = string(blackfriday.MarkdownCommon([]byte(body)))
	}
	return body
}
