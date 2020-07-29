package src

import (
	"fmt"
	"github.com/DumkaUA/dumka_go/src/model"
	"github.com/jinzhu/gorm"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"strconv"
	"time"
)

func RunAdminBot(db *gorm.DB) {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_ADMIN_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(b.Me.Username)

	b.Handle("/start", func(m *tb.Message) {
		_, _ = b.Send(m.Sender, "Hello World!")

		var schools []model.School

		db.Find(&schools)

		var response = ""
		for _, school := range schools {
			response += "School #" + strconv.FormatInt(school.Id, 10) + ": " + school.Name + "\n"
		}

		_, _ = b.Send(m.Sender, response)

	})

	b.Start()
}
