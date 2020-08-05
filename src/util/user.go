package util

import (
	"github.com/DumkaUA/dumka_go/src/model"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) *model.User {
	user, err := c.Keys["USER"].(*model.User)

	if !err {
		panic("ERROR: NO USER IN THE CONTEXT")
	}

	return user
}
