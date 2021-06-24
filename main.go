package main

import (
	"log"

	"internal/twitter"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	quote := GetLeastRecentPostedQuote()
	status := quote.text + " — " + quote.author

	twitter.PostTweetStatusUpdate(twitter.TwitterClient(), status)
	UpdatePostDate(quote.Id)
}
