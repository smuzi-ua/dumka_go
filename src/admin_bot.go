package src

import (
	"fmt"
	"github.com/DumkaUA/dumka_go/src/model"
	"github.com/jinzhu/gorm"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"strconv"
	"strings"
	"time"
)

func isAdmin(username string, db *gorm.DB) bool { //checks whether a bot user has rights to use it
	var admins []model.Admin
	db.Find(&admins)
	for i := 0; i < len(admins); i++ {
		if username == admins[i].TelegramNickname {
			return true
		}
	}
	return false
}

func RunAdminBot(db *gorm.DB) {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_ADMIN_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(b.Me.Username)

	b.Handle("/add_admin", func(m *tb.Message) { // command: /add_admin telegram_nickname
		if isAdmin(m.Sender.Username, db) {
			var admin = model.Admin{
				//what should I do with ID?
				TelegramNickname: strings.Split(m.Text, " ")[1],
			}
			db.Create(&admin)
		}
		b.Send(m.Sender, "Admin successfully added")
	})

	b.Handle("/remove_admin", func(m *tb.Message) {
		if isAdmin(m.Sender.Username, db) {
			// todo later
			// will remove row from 'admins' table
		}
	})

	b.Handle("/add_school", func(m *tb.Message) { //command: /add_school name
		if isAdmin(m.Sender.Username, db) {
			var school = model.School{
				//what should I do with ID?
				Name:    strings.Split(m.Text, " ")[1],
				Display: false, // what should I do with "Display"?
			}
			db.Create(&school)
			b.Send(m.Sender, "School successfully added")
		}
	})

	b.Handle("/add_user", func(m *tb.Message) {
		if isAdmin(m.Sender.Username, db) {
			// todo later
		}
	})

	b.Handle("/get_users", func(m *tb.Message) { //command: /get_users
		if isAdmin(m.Sender.Username, db) {
			var users []model.User

			db.Find(&users)

			for _, user := range users {
				var response = "User #" + strconv.Itoa(user.Id) + ": " + user.Name + "\n" + "User's nickname: " + user.Nickname + "\n"

				_, _ = b.Send(m.Sender, response)
			}
		}
	})

	b.Handle("/get_users_info", func(m *tb.Message) { //command: /get_users_info
		if isAdmin(m.Sender.Username, db) {
			var users []model.User

			db.Find(&users)

			for _, user := range users {
				var response = "User #" + strconv.Itoa(user.Id) + ": " + user.Name + "\n" + "User's nickname: " + user.Nickname + "\n"

				var proposals []model.Proposal
				db.Where(&model.Proposal{UserId: user.Id}).Find(&proposals)
				response += "Number of proposals posted by user: " + strconv.Itoa(len(proposals)) + "\n"

				var comments []model.Comment
				db.Where(&model.Comment{UserId: user.Id}).Find(&comments)
				response += "Number of comments written by user: " + strconv.Itoa(len(comments)) + "\n"

				var proposal_likes []model.ProposalLike
				db.Where(&model.ProposalLike{UserId: user.Id}).Find(&proposal_likes)

				var comment_likes []model.CommentLike
				db.Where(&model.CommentLike{UserId: user.Id}).Find(&comment_likes)

				response += "Number of proposals/comments student liked: " + strconv.Itoa(len(proposal_likes)+len(comment_likes)) + "\n"

				_, _ = b.Send(m.Sender, response)
			}

		}
	})

	b.Handle("/get_user", func(m *tb.Message) {
		if isAdmin(m.Sender.Username, db) {

		}
	})

	b.Handle("/get_schools", func(m *tb.Message) {
		if isAdmin(m.Sender.Username, db) {
			var schools []model.School

			db.Find(&schools)

			var response = ""
			for _, school := range schools {
				response += "School #" + strconv.Itoa(school.Id) + ": " + school.Name + "\n"
			}

			_, _ = b.Send(m.Sender, response)
		}
	})

	b.Handle("/stats", func(m *tb.Message) {
		if isAdmin(m.Sender.Username, db) {
			// todo later
		}
	})

	b.Start()
}
