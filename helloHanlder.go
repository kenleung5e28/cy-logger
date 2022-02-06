package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
)

func helloHandler(b *tb.Bot, m *tb.Message) {
	_, err := b.Send(m.Sender, fmt.Sprintf("Hello World! Chat ID: %d", m.Chat.ID))
	if err != nil {
		log.Println("Error sending message in /hello handle: ", err)
		return
	}
}
