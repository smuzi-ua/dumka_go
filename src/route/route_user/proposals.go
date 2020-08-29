package route_user

import (
	"fmt"
	"github.com/DumkaUA/dumka_go/src/model"
	"github.com/DumkaUA/dumka_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// todo archived proposals
func ProposalsRoute(c *gin.Context) {
	u := util.GetUser(c)
	db := util.GetDB(c)

	var data []model.Proposal
	var displayData []model.Proposal

	stages := "('active', 'in_progress', 'archive')"

	if u.Type == "supervisor" {
		stages = "('active', 'in_progress', 'archive', 'none')"
	}

	// todo there is probably a better way to do that
	fmt.Println(db.Where("user_id IN (SELECT id FROM users WHERE school_id = ?) AND stage IN "+stages, u.SchoolId).Find(&data).Error)

	for _, r := range data {
		if r.Anonymous {
			r.UserId = 0
			displayData = append(displayData, r)
		} else {
			displayData = append(displayData, r)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": displayData,
	})
}
