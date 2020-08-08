package route

import (
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

	// todo there is probably a better way to do that
	db.Where("user_id IN (SELECT id FROM users WHERE school_id = ?)", u.SchoolId).Find(&data)

	for _, r := range data {
		if r.Anonymous {
			displayData = append(displayData, model.Proposal{Id: r.Id, Title: r.Title, Text: r.Text, Date: r.Date, Anonymous: r.Anonymous})
		} else {
			displayData = append(displayData, model.Proposal{Id: r.Id, Title: r.Title, Text: r.Text, Date: r.Date, Anonymous: r.Anonymous, UserId: r.UserId})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": displayData,
	})
}
