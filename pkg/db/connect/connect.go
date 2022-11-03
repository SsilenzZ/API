package connect

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func init() {
	var err error
	if os.Getenv("DB_CONNECTION") == "" {
		err = os.Setenv("DB_CONNECTION", "root:@tcp(127.0.0.1:3306)/commentsdb?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			panic(err)
		}
	}
	dsn := os.Getenv("DB_CONNECTION")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{DisableAutomaticPing: false})
	if err != nil {
		panic(err)
	}
}
