package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

type Quote struct {
	Id     int
	text   string
	author string
}

func GetLeastRecentPostedQuote() *Quote {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Panicf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	query := "select id, text, author from quotes order by last_posted_at DESC LIMIT 1"
	quote := &Quote{}
	err = conn.QueryRow(context.Background(), query).Scan(&quote.Id, &quote.text, &quote.author)

	if err != nil {
		log.Panicf("Unable to connect to database: %v\n", err)
	}

	log.Printf("Quote %v", quote)

	return quote
}
