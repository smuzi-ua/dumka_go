package util

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetDB(c *gin.Context) *gorm.DB {
	db, err := c.Keys["DB"].(*gorm.DB)

	if !err {
		panic("ERROR: NO DATABASE IN THE CONTEXT")
	}

	return db
}
