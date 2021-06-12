package glob

import (
	"os"
)

var WorkDir string

func init() {
	WorkDir, _ = os.Getwd()
}
