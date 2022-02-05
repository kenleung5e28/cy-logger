package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func isAllowedUser(chatId int64) bool {
	myChatId, err := strconv.ParseInt(os.Getenv("MY_CHAT_ID"), 10, 64)
	return err != nil || chatId == myChatId
}

func main() {
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 20 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		if !isAllowedUser(m.Chat.ID) {
			_, err := b.Send(m.Sender, "You are not authorized, sorry.")
			if err != nil {
				log.Println("Error sending unauthorized message: ", err)
			}
			return
		}
		_, err := b.Send(m.Sender, fmt.Sprintf("Hello World! Chat ID: %d", m.Chat.ID))
		if err != nil {
			log.Println("Error sending message in /hello handle: ", err)
			return
		}
	})

	b.Start()
}
