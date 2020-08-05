package api

import (
	"github.com/DumkaUA/dumka_go/src/model"
	"github.com/DumkaUA/dumka_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func StatusRoute(c *gin.Context) {
	var schools uint
	var users uint

	db := util.GetDB(c)

	db.Table("schools").Where(model.School{Display: true}).Count(&schools)
	db.Table("users").Where(model.User{Verified: true}).Count(&users)

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
		"data": gin.H{
			"time":    time.Now(),
			"schools": schools,
			"users":   users,
		},
	})
}
