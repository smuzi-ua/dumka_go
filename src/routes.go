package src

import (
	"github.com/DumkaUA/dumka_go/src/api"
	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {

	v1 := server.Group("/v1")

	{
		v1.POST("/status", api.StatusRoute)
		v1.POST("/schools", api.SchoolsRoute)
		v1.POST("/report_categories", api.ReportCategoriesRoute)

		v1.POST("/auth", api.AuthRoute)
	}

}
