package testutil

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

func NewTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx, recorder
}

func SetAuthUser(c *gin.Context, userID string) {
	c.Set("userID", userID)
}

func MakeGet(c *gin.Context, params gin.Params, query url.Values) {
	c.Request.Method = http.MethodGet
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = params
	c.Request.URL.RawQuery = query.Encode()
}

func MakePost(c *gin.Context, params gin.Params, body any) {
	b, _ := json.Marshal(body)

	c.Request.Method = http.MethodPost
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = params
	c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
}

func MakePatch(c *gin.Context, params gin.Params, body any) {
	b, _ := json.Marshal(body)

	c.Request.Method = http.MethodPatch
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = params
	c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
}

func MakeDelete(c *gin.Context, params gin.Params) {
	c.Request.Method = http.MethodDelete
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = params
}
