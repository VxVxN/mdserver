package glob

import (
	"os"
)

var (
	WorkDir      string
	SessionToken string // TODO: Redo it
)

func init() {
	WorkDir, _ = os.Getwd()
}
