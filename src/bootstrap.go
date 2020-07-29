package src

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Bootstrap(server *gin.Engine, db *gorm.DB) {
	// making database available for routes
	server.Use(func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	})

	Routes(server)

	RunAdminBot(db)
}
