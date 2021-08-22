package tools

import (
	"path"

	"github.com/VxVxN/mdserver/pkg/consts"
)

func GetPathToTMPImages(username string) string {
	return path.Join(consts.PathToPosts, username, "tmp_images")
}
