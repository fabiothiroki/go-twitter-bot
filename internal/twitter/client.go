package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func TwitterClient() *twitter.Client {
	creds := GetTwitterCreds()
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}
