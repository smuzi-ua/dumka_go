package route_user

import (
	"github.com/DumkaUA/dumka_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRoute(c *gin.Context) {
	u := util.GetUser(c)
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
		"data": gin.H{
			"name":     u.Name,
			"nickname": u.Nickname,
			"school":   u.SchoolId,
			"type":     u.Type,
		},
	})
}
