package main

import (
	_ "Api/config"
	"Api/pkg/handler"
	"Api/pkg/router"
	"Api/pkg/service"
	"os"
)

func main() {
	authHandler := handler.NewAuthHandler(service.BcryptHasher{Cost: 5}, service.Jwt{})
	e := router.New(&authHandler)
	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}
