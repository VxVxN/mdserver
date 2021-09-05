package mock

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	gSessions "github.com/gorilla/sessions"
)

type Store struct{}

func (store *Store) Options(sessions.Options) {
}

func (store *Store) New(r *http.Request, name string) (*gSessions.Session, error) {
	return new(gSessions.Session), nil
}

func (store *Store) Save(r *http.Request, w http.ResponseWriter, s *gSessions.Session) error {
	return nil
}

func (store *Store) Get(r *http.Request, name string) (*gSessions.Session, error) {
	var s gSessions.Session

	s.Values = make(map[interface{}]interface{})
	s.Values["username"] = "test"

	return &s, nil
}
