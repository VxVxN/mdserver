package post

import (
	"os"
	"path"
	"testing"

	"github.com/VxVxN/mdserver/internal/glob"
	"github.com/VxVxN/mdserver/pkg/consts"
	"github.com/stretchr/testify/require"
)

func TestSavePost(t *testing.T) {
	dirName := "test_dir"
	fileName := "test_post"
	text := "test text"

	var err error
	glob.WorkDir, err = os.Getwd()
	require.NoError(t, err)
	glob.WorkDir = path.Dir(path.Dir(glob.WorkDir))
	consts.PathToPosts = path.Join(glob.WorkDir, "..", "posts")

	pathToFile := path.Join(glob.WorkDir, "..", "posts", dirName, fileName) + consts.ExtMd
	pathToDir := path.Join(glob.WorkDir, "..", "posts", dirName)

	err = os.Mkdir(pathToDir, 0777)
	require.NoError(t, err)

	defer os.RemoveAll(pathToDir)

	// check that it will not save the post without a file
	errObj := SavePost(dirName, fileName, text, false)
	require.NotNil(t, errObj, nil)

	err = os.WriteFile(pathToFile, []byte("text"), 0644)
	require.NoError(t, err)

	// check that it will save the post
	errObj = SavePost(dirName, fileName, text, false)
	require.Nil(t, errObj)

	data, err := os.ReadFile(pathToFile)
	require.NoError(t, err)
	require.Equal(t, string(data), text)

	err = os.Remove(pathToFile)
	require.NoError(t, err)

	// check that it will create file
	errObj = SavePost(dirName, fileName, text, true)
	require.Nil(t, errObj)

	data, err = os.ReadFile(pathToFile)
	require.NoError(t, err)
	require.Equal(t, string(data), text)
}
