package router

import (
	"Api/pkg/handler"
	_ "Api/pkg/router/docs"
	"Api/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

// @title Echo Swagger API
// @version 1.0
// @description This is an echo swagger API server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func New(authHandler *handler.AuthHandler, commentHandler *handler.CommentHandler, postHandler *handler.PostHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/comments", commentHandler.ReturnAllComments)
	e.GET("/comments/:id", commentHandler.ReturnComment)
	e.GET("/posts", postHandler.ReturnAllPosts)
	e.GET("/posts/:id", postHandler.ReturnPost)
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
	a.POST("/comments", commentHandler.CreateComment)
	a.DELETE("/comments/:id", commentHandler.DeleteComment)
	a.PUT("/comments/:id", commentHandler.UpdateComment)
	a.POST("/posts", postHandler.CreatePost)
	a.DELETE("/posts/:id", postHandler.DeletePost)
	a.PUT("/posts/:id", postHandler.UpdatePost)

	return e
}
