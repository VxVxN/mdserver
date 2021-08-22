package consts

import (
	"path"

	"github.com/VxVxN/mdserver/internal/glob"
)

const (
	ExtMd = ".md"
)

var (
	PathToTemplates    = path.Join(glob.WorkDir, "templates")
	PathToPosts        = path.Join(glob.WorkDir, "posts")
	PathToImages       = path.Join(PathToPosts, "images")
	PathToStatic       = path.Join(glob.WorkDir + "/public/static")
	PathToStaticImages = path.Join(PathToStatic, "images")
)
