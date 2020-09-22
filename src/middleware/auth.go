package middleware

import (
	"github.com/DumkaUA/dumka_go/src/model"
	"github.com/DumkaUA/dumka_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// todo use a library for auth

		var q struct {
			Authorization string `form:"Authorization" binding:"required"`
		}

		if err := c.ShouldBindHeader(&q); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"ok":         false,
				"error":      err.Error(),
				"error_code": -2,
			})
			return
		}

		if len(strings.Split(q.Authorization, " ")) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"ok":         false,
				"error":      "Weird Authorization parameter :/",
				"error_code": -2,
			})
			return
		}

		// todo check for sql injection (?)
		token := strings.Split(q.Authorization, " ")[1]

		db := util.GetDB(c)

		var user model.User
		db.Where(model.User{Token: token}).First(&user)

		if user.Id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"ok":         false,
				"error":      "USER WITH SUCH TOKEN DOES NOT EXIST",
				"error_code": -2,
			})
			return
		}

		if !user.Verified {
			c.JSON(http.StatusBadRequest, gin.H{
				"ok":         false,
				"error":      "USER'S ACCOUNT IS NOT VERIFIED",
				"error_code": -1,
			})
			return
		}

		c.Set("USER", &user)

		c.Next()
	}
}
