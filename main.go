package main

import (
	"github.com/DumkaUA/dumka_go/src"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// todo tests
func main() {
	db, err := gorm.Open("mysql", "root:@/main?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	server := gin.Default()

	// making database available for every gin route
	server.Use(func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	})

	// this is where all the magic happens
	src.Bootstrap(server)

	err = server.Run()

	if err != nil {
		panic(err)
	}
}
