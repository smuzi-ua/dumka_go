package route_user

import (
	"fmt"
	"github.com/DumkaUA/dumka_go/src/model"
	"github.com/DumkaUA/dumka_go/src/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func ProposalsAddRoute(c *gin.Context) {
	u := util.GetUser(c)
	db := util.GetDB(c)

	var q struct {
		Title     string `form:"title" binding:"required"`
		Text      string `form:"text" binding:"required"`
		Deadline  string `form:"deadline" binding:"required"`
		Anonymous bool   `form:"anonymous"`
	}

	if err := c.ShouldBindBodyWith(&q, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":         false,
			"error":      err.Error(),
			"error_code": -1,
		})
		return
	}
	fmt.Println(q)

	if q.Deadline != "day" && q.Deadline != "2day" && q.Deadline != "week" && q.Deadline != "month" {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":         false,
			"error":      "invalid deadline",
			"error_code": 1,
		})
		return
	}

	if len(q.Title) > 512 {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":         false,
			"error":      "invalid title",
			"error_code": 2,
		})
		return
	}

	db.Create(&model.Proposal{UserId: u.Id, Anonymous: q.Anonymous, Title: q.Title, Text: q.Text, Deadline: q.Deadline})

	// todo organize all error codes properly

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
