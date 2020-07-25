package main

import (
	"github.com/DumkaUA/dumka_go/src"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rushteam/gosql"

	"github.com/gin-gonic/gin"
)

func main() {
	db := gosql.NewCluster(
		gosql.AddDb("mysql", "root:@tcp(127.0.0.1:3306)/main?parseTime=true&readTimeout=3s&writeTimeout=3s&timeout=3s"),
	)

	server := gin.Default()

	server.Use(func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	})

	// this is where all the magic happens
	src.Bootstrap(server)

	err := server.Run()

	if err != nil {
		panic(err)
	}
}
