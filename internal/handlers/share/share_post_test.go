package share

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"github.com/VxVxN/mdserver/internal/driver/mongo/mocks"
	"github.com/VxVxN/mdserver/pkg/mock"
)

func TestController_SharePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := NewController(new(mocks.MongoShare))

	var store mock.Store
	r := gin.Default()

	r.Use(sessions.Sessions("test", &store))
	r.POST("share_link", ctrl.SharePostHandler)

	request := RequestSharePost{
		DirName:  "test_dir",
		FileName: "test_file",
	}

	requestBytes, err := json.Marshal(request)
	require.NoError(t, err, "cannot marshal request")

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/share_link", bytes.NewReader(requestBytes))

	r.ServeHTTP(res, req)

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(res.Result().Body)

	resp := make(map[string]interface{})
	_ = json.Unmarshal(buf.Bytes(), &resp)
	link, _ := resp["link"].(string)

	require.Equal(t, mocks.MockGeneratedLink, link, "Not expected link")
}
