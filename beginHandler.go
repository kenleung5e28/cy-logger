package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
)

func beginHandler(b *tb.Bot, m *tb.Message) {
	// TODO
}

func saveRecord(chatId int64, description string) error {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

}
