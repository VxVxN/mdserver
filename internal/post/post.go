package post

import (
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
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
		// не файл, а папка
		return post{}, 404, fmt.Errorf("dir")
	}
	val, ok := p.Items[md+strconv.FormatBool(isEdit)]
	if !ok || (ok && val.ModTime != info.ModTime().UnixNano()) {
		p.RLock()
		defer p.RUnlock()
		fileText, _ := ioutil.ReadFile(md)
		lines := strings.Split(string(fileText), "\n")
		title := lines[0]
		var body string
		if isEdit {
			body = "<textarea >" + strings.Join(lines[1:], "\n") + "</textarea >"
		} else {
			body = strings.Join(lines[1:], "\n")
			body = string(blackfriday.MarkdownCommon([]byte(body)))
		}
		p.Items[md] = post{title, template.HTML(body), info.ModTime().UnixNano()}
	}
	post := p.Items[md]
	return post, 200, nil
}
