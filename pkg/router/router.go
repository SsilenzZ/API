package router

import (
	"Api/pkg/handler"
	"Api/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func New(authHandler *handler.AuthHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/", handler.HelloWorld)
	e.GET("/comments", handler.ReturnAllComments)
	e.GET("/comments/:id", handler.ReturnComment)
	e.GET("/posts", handler.ReturnAllPosts)
	e.GET("/posts/:id", handler.ReturnPost)
	e.POST("/sign-up", authHandler.SignUp)
	e.POST("/sign-in", authHandler.SignIn)
	e.POST("/sign-in/google", authHandler.GoogleSignIn)
	e.GET("token/google", authHandler.GetAuthToken)

	a := e.Group("/api")
	config := middleware.JWTConfig{
		Claims:     &service.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}

	a.Use(middleware.JWTWithConfig(config))
	a.POST("/comments", handler.CreateComment)
	a.DELETE("/comments/:id", handler.DeleteComment)
	a.PUT("/comments/:id", handler.UpdateComment)
	a.POST("/posts", handler.CreatePost)
	a.DELETE("/posts/:id", handler.DeletePost)
	a.PUT("/posts/:id", handler.UpdatePost)

	return e
}
