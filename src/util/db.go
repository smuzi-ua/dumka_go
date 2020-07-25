package util

import (
	"github.com/gin-gonic/gin"
	"github.com/rushteam/gosql"
)

func GetDB(c *gin.Context) *gosql.PoolCluster {
	db, err := c.Keys["DB"].(*gosql.PoolCluster)

	if !err {
		panic("ERROR: NO DATABASE IN THE CONTEXT")
	}

	return db
}
