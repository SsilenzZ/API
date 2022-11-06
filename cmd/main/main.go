package main

import (
	_ "Api/config"
	"Api/pkg/handler"
	"Api/pkg/router"
	"Api/pkg/service"
	"os"
)

func main() {
	authHandler := handler.NewAuthHandler(service.BcryptHasher{Cost: 5}, service.Jwt{}, service.Oauth{})
	commentHandler := handler.CommentHandler{}
	postHandler := handler.PostHandler{}
	e := router.New(&authHandler, &commentHandler, &postHandler)

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}
