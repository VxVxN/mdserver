package glob

import (
	"os"
	"path"
)

var WorkDir string

func init() {
	WorkDir = path.Dir(os.Args[0])
}
