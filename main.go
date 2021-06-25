package main

import (
	"log"

	"internal/database"
	"internal/twitter"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	quote := database.GetLeastRecentPostedQuote()
	status := quote.Text + " — " + quote.Author

	twitter.PostTweetStatusUpdate(twitter.TwitterClient(), status)
	database.UpdatePostDate(quote.ID)
}
