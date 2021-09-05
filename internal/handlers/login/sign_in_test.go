package login

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
	tools "github.com/VxVxN/mdserver/pkg/tools/test"
)

func TestController_SignIn(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := NewController(new(mocks.MongoUsers))

	store := tools.NewStore()

	r := gin.Default()
	r.Use(sessions.Sessions("test", store))

	r.POST("sign_in", ctrl.SignIn)

	expectedUsername := "test"
	request := RequestSignIn{
		Username: expectedUsername,
		Password: "123",
	}

	requestBytes, err := json.Marshal(request)
	require.NoError(t, err, "cannot marshal request")

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sign_in", bytes.NewReader(requestBytes))
	r.ServeHTTP(res, req)
	session, _ := store.Get(req, "test")
	username, _ := session.Values["username"]

	require.Equal(t, expectedUsername, username, "Incorrect username saved in session")
}
