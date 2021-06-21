package main

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	creds := getCreds()
	client := getClient(creds)

	tweet, resp, err := client.Statuses.Update("just setting up my twttr", nil)

	if err != nil {
		log.Fatal("Error publishing tweet", err)
	}

	fmt.Println("Hello World! ", resp, tweet)
}

func getClient(creds *creds) *twitter.Client {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}
