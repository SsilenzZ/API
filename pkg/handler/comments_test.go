package handler

import (
	_ "Api/config"
	"Api/pkg/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var (
	commentJSON = `{"Post":1,"ID":1,"Name":"test","Email":"test@mail.com","Body":"some body"}
`
	commentsJSON = `[{"Post":1,"ID":1,"Name":"test","Email":"test@mail.com","Body":"some body"},` +
		`{"Post":1,"ID":2,"Name":"test 2","Email":"test@mail.com","Body":"some body"},` +
		`{"Post":1,"ID":3,"Name":"test 3","Email":"test@mail.com","Body":"some another body"}]
`
	createCommentJSON, _ = json.Marshal(map[string]interface{}{"Post": 1, "ID": 4, "Name": "create comment test", "Body": "some body"})
	updateCommentJSON, _ = json.Marshal(map[string]interface{}{"Name": "update comment test", "Body": "some edited body"})
	commentHandler       = NewCommentHandler()
)

func TestReturnAllComments(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/comments")

	if assert.NoError(t, commentHandler.ReturnAllComments(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, commentsJSON, rec.Body.String())
	}
}

func TestReturnComment(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/comments/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, commentHandler.ReturnComment(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, commentJSON, rec.Body.String())
	}
}

func TestCreateComment(t *testing.T) {
	e := echo.New()

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}

	h := middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &service.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	})(handler)

	token := os.Getenv("TOKEN_FOR_TEST")

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(createCommentJSON))
	res := httptest.NewRecorder()
	req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+token)
	c := e.NewContext(req, res)
	c.SetPath("/api/comments")
	assert.NoError(t, h(c))

	if assert.NoError(t, commentHandler.CreateComment(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
	}
}

func TestUpdateComment(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(updateCommentJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/comments/:id")
	c.SetParamNames("id")
	c.SetParamValues("4")

	if assert.NoError(t, commentHandler.UpdateComment(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteComment(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/comments/:id")
	c.SetParamNames("id")
	c.SetParamValues("4")

	if assert.NoError(t, commentHandler.DeleteComment(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
