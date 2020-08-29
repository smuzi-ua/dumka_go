package main

import (
	"github.com/DumkaUA/dumka_go/src"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"math/rand"
	"os"
	"time"
)

// todo tests
// todo organize all errors in a nice files and give them unique id
func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	db, err := gorm.Open("mysql", os.Getenv("DB_URL"))

	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())

	server := gin.Default()

	// this is where all the magic happens
	src.Bootstrap(server, db)

	err = server.Run()

	if err != nil {
		panic(err)
	}
}
