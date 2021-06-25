package main

import (
	"log"

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

	quote := database.GetLeastRecentPostedQuote()
	status := postformatter.Format(quote)

	twitter.PostTweetStatusUpdate(twitter.TwitterClient(), status)
	database.UpdatePostDate(quote.ID)
}
