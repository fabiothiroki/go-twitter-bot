package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	quote := GetLeastRecentPostedQuote()

	PostTweetStatusUpdate(quote.text + " — " + quote.author)
	UpdatePostDate(quote.Id)
}
