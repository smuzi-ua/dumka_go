package route

import (
	"github.com/DumkaUA/dumka_go/src/model"
	"github.com/DumkaUA/dumka_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReportCategoriesRoute(c *gin.Context) {
	var categories []model.ReportCategory
	db := util.GetDB(c)
	db.Find(&categories)
	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": categories,
	})
}
