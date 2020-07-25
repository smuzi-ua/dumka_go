package src

import (
	"github.com/gin-gonic/gin"
)

func Bootstrap(server *gin.Engine) {
	Routes(server)
}
