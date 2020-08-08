package route

import (
	"fmt"
	"github.com/DumkaUA/dumka_go/src/model"
	"github.com/DumkaUA/dumka_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// todo consider bruteforce

func AuthRoute(c *gin.Context) {

	var q struct {
		School   int    `form:"school" binding:"required"`
		Nickname string `form:"nickname" binding:"required"`
		Name     string `form:"name" binding:"required"`
		Code     int    `form:"code"`
	}

	if err := c.ShouldBind(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":         false,
			"error":      err.Error(),
			"error_code": -1,
		})
		return
	}

	db := util.GetDB(c)

	fmt.Println(model.School{}.TableName())
	// validating School
	var schoolCount int
	db.Table(model.School{}.TableName()).Where(model.School{Id: q.School}).Count(&schoolCount)
	if schoolCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":         false,
			"error":      "SCHOOL DOES NOT EXIST",
			"error_code": -1,
		})
		return
	}

	if q.Code == 0 {

		var userCount int
		db.Table(model.User{}.TableName()).Where(model.User{SchoolId: q.School, Nickname: q.Nickname}).Count(&userCount)

		if userCount == 0 {
			db.Create(&model.User{
				SchoolId: q.School,
				Name:     q.Name,
				Nickname: q.Nickname,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})

		return
	}

	var user model.User

	db.Where(model.User{SchoolId: q.School, Nickname: q.Nickname, Name: q.Name}).Take(&user)

	// if user doesn't exist
	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":         false,
			"error":      "USER DOES NOT EXIST",
			"error_code": -1,
		})
		return
	}

	if !user.Verified {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":         false,
			"error":      "USER IS NOT VERIFIED",
			"error_code": 1,
		})
		return
	}

	if user.Code != q.Code {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":         false,
			"error":      "CODE DOES NOT MATCH",
			"error_code": 2,
		})
		return
	}

	user.UpdateToken()
	user.UpdateCode()
	db.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"ok":        true,
		"user_type": user.Type,
		"token":     user.Token,
	})
}
