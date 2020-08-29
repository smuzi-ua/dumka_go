package route_user

import (
	"github.com/DumkaUA/dumka_go/src/model"
	"github.com/DumkaUA/dumka_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SchoolRoute(c *gin.Context) {
	u := util.GetUser(c)
	db := util.GetDB(c)

	var data []model.User
	var school model.School

	db.Where(model.User{SchoolId: u.SchoolId, Verified: true}).Select([]string{"id", "name", "nickname"}).Find(&data)
	db.Where(model.School{Id: u.SchoolId}).Select([]string{"name"}).First(&school)

	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": data,
		"name": school.Name,
	})
}
