package main

import (
	"context"
	"log"
	"time"

	"github.com/fabiothiroki/go-twitter-bot/internal/postformatter"
	"github.com/fabiothiroki/go-twitter-bot/internal/twitter"

	"github.com/fabiothiroki/go-twitter-bot/internal/database"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	conn := database.OpenConnection()
	defer conn.Close(context.Background())
	dbService := &database.DbService{DbConn: conn}

	quote := dbService.GetLeastRecentPostedQuote()
	status := postformatter.Format(quote)

	twitter.PostTweetStatusUpdate(twitter.TwitterClient(), status)
	dbService.UpdatePostDate(quote.ID, time.Now())
}
