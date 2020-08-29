package src

import (
	"github.com/DumkaUA/dumka_go/src/middleware"
	"github.com/DumkaUA/dumka_go/src/route/route_open"
	"github.com/DumkaUA/dumka_go/src/route/route_user"
	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {

	v1 := server.Group("/v1")

	v1s := v1.Group("/u")
	v1s.Use(middleware.Auth())

	{
		v1.POST("/status", route_open.StatusRoute)
		v1.POST("/schools", route_open.SchoolsRoute)
		v1.POST("/report_categories", route_open.ReportCategoriesRoute)
		v1.POST("/auth", route_open.AuthRoute)

		v1s.POST("/user", route_user.UserRoute)
		v1s.POST("/proposals", route_user.ProposalsRoute)
		v1s.POST("/proposals_add", route_user.ProposalsAddRoute)
		v1s.POST("/school", route_user.SchoolRoute)
	}
}
