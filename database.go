package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

type Quote struct {
	Id     int
	text   string
	author string
}

func openConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Panicf("Unable to connect to database: %v\n", err)
	}

	return conn
}

func GetLeastRecentPostedQuote() *Quote {
	conn := openConnection()
	defer conn.Close(context.Background())

	query := "select id, text, author from quotes order by last_posted_at DESC LIMIT 1"
	quote := &Quote{}
	conn.QueryRow(context.Background(), query).Scan(&quote.Id, &quote.text, &quote.author)

	log.Printf("Quote %v", quote)

	return quote
}

func UpdatePostDate(quoteId int) {
	conn := openConnection()
	defer conn.Close(context.Background())

	query := "update quotes set last_posted_at=$1 where id=$2"
	conn.QueryRow(context.Background(), query, time.Now(), quoteId)
}
