package route_open

import (
	"github.com/DumkaUA/dumka_go/src/model"
	"github.com/DumkaUA/dumka_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SchoolsRoute(c *gin.Context) {
	var schools []model.School
	db := util.GetDB(c)
	db.Where(&model.School{Display: true}).Find(&schools)
	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": schools,
	})
}
