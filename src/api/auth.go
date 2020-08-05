package api

import (
	"github.com/DumkaUA/dumka_go/src/model"
	"github.com/DumkaUA/dumka_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// todo consider bruteforce
// todo use nickname instead of real name (to avoid repetition)

func AuthRoute(c *gin.Context) {

	var q struct {
		School int    `form:"school" binding:"required"`
		Name   string `form:"name" binding:"required"`
		Code   int    `form:"code"`
	}

	if err := c.ShouldBind(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
		return
	}

	db := util.GetDB(c)

	// validating School
	var schoolCount int
	db.Table("schools").Where(model.School{Id: q.School}).Count(&schoolCount)

	if schoolCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "SCHOOL DOES NOT EXIST"})
		return
	}

	if q.Code == 0 {

		var userCount int
		db.Table("users").Where(model.User{SchoolId: q.School, Name: q.Name}).Count(&userCount)

		if userCount == 0 {
			db.Create(model.User{
				SchoolId: q.School,
				Name:     q.Name,
				Type:     "student",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})

		return
	}

	// todo AUTH

}
