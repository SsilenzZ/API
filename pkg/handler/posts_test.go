package handler

import (
	_ "Api/config"
	"Api/pkg/service"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	stringPostsJSON = `[{"ID":1,"User":1,"Title":"test","Body":"some body"},` +
		`{"ID":2,"User":1,"Title":"test 2","Body":"some body"}]
`
	stringPostJSON = `{"ID":1,"User":1,"Title":"test","Body":"some body"}
`
	createPostJSON, _ = json.Marshal(map[string]interface{}{"ID": 3, "Title": "create post test", "Body": "some body"})
	updatePostJSON, _ = json.Marshal(map[string]interface{}{"Title": "update post test", "Body": "some edited body"})
	postHandler       = NewPostHandler()
)

func TestReturnAllPosts(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts")

	if assert.NoError(t, postHandler.ReturnAllPosts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, stringPostsJSON, rec.Body.String())
	}
}

func TestReturnPost(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, postHandler.ReturnPost(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, stringPostJSON, rec.Body.String())
	}
}

func TestCreatePost(t *testing.T) {
	e := echo.New()

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}

	h := middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &service.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	})(handler)

	token := os.Getenv("TOKEN_FOR_TEST")

	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(createPostJSON))
	res := httptest.NewRecorder()
	req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+token)
	c := e.NewContext(req, res)
	c.SetPath("/api/posts")
	assert.NoError(t, h(c))

	if assert.NoError(t, postHandler.CreatePost(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
	}
}

func TestUpdatePost(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(updatePostJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("3")

	if assert.NoError(t, postHandler.UpdatePost(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeletePost(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("3")

	if assert.NoError(t, postHandler.DeletePost(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
