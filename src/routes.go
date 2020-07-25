package src

import (
	"github.com/DumkaUA/dumka_go/src/api"
	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {
	server.GET("/schools", api.SchoolsRoute)
}
