package api

import (
	"github.com/DumkaUA/dumka_go/src/model"
	"github.com/DumkaUA/dumka_go/src/util"
	"github.com/gin-gonic/gin"
	"github.com/rushteam/gosql"
	"net/http"
)

func SchoolsRoute(c *gin.Context) {
	var schools []model.School

	db := util.GetDB(c)

	err := db.FetchAll(&schools,
		gosql.Columns("id", "name"),
	)

	if err != nil {
		c.JSON(200, gin.H{"data": nil})
	}

	c.JSON(http.StatusOK, gin.H{"data": schools})
	//c.String(200, "Success")
}
