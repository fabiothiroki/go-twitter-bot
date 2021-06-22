package main

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func PostTweetStatusUpdate(status string) {
	creds := getCreds()
	client := getClient(creds)

	_, resp, err := client.Statuses.Update(status, nil)

	if err != nil {
		log.Fatal("Error publishing tweet", err)
	}

	log.Println("Post tweet response", resp.Status)

}

func getClient(creds *creds) *twitter.Client {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}
