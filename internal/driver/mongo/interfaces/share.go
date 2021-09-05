package interfaces

import (
	"github.com/VxVxN/mdserver/internal/driver/mongo/share"
	e "github.com/VxVxN/mdserver/pkg/error"
)

type MongoShare interface {
	GenerateLink(username, dirName, fileName string) (string, *e.ErrObject)
	GetLinks(username string) (*share.Share, error)
	DeleteLink(username, link string) *e.ErrObject
}
