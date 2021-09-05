package share

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"github.com/VxVxN/mdserver/internal/driver/mongo/mocks"
	"github.com/VxVxN/mdserver/internal/glob"
	"github.com/VxVxN/mdserver/pkg/config"
	"github.com/VxVxN/mdserver/pkg/mock"
)

func TestController_ShareLinksHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.InitTestConfig()

	ctrl := NewController(new(mocks.MongoShare))

	var store mock.Store
	r := gin.Default()

	r.Use(sessions.Sessions("test", &store))
	r.LoadHTMLGlob(glob.WorkDir + "/../../../templates/*")
	r.GET("share_links", ctrl.ShareLinksHandler)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/share_links", nil)

	r.ServeHTTP(res, req)

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(res.Result().Body)

	expectedLinks := []string{
		"https://testDomain.com/share/test/6ac6d7ef-b7cd-487a-9cf5-9e2d6fe5397f",
		"https://testDomain.com/share/test/57afcf14-0b2b-4a1e-be73-d16fbd5df032",
		"https://testDomain.com/share/test/12e446b7-ee8a-41ef-8f7c-a951968090b3",
	}
	for _, link := range expectedLinks {
		if !strings.Contains(buf.String(), link) {
			require.Fail(t, "Not found link in html", link)
		}
	}
}
