package src

import (
	"github.com/DumkaUA/dumka_go/src/middleware"
	"github.com/DumkaUA/dumka_go/src/route"
	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {

	v1 := server.Group("/v1")

	v1s := v1.Group("/u")
	v1s.Use(middleware.Auth())

	{
		v1.POST("/status", route.StatusRoute)
		v1.POST("/schools", route.SchoolsRoute)
		v1.POST("/report_categories", route.ReportCategoriesRoute)
		v1.POST("/auth", route.AuthRoute)

		v1s.POST("/user", route.UserRoute)
		v1s.POST("/proposals", route.ProposalsRoute)
		v1s.POST("/school", route.SchoolRoute)

	}
}
