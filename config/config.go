package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

const projectDirName = "Api"

func init() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Print("Error loading .env file")
	}
}
