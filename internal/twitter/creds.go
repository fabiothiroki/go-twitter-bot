package twitter

import (
	"os"
)

// Creds represents the values required for twitter authentication
type Creds struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

// GetTwitterCreds returns the keys required for authentication
func GetTwitterCreds() *Creds {
	return &Creds{
		ConsumerKey:    os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
		AccessToken:    os.Getenv("ACCESS_TOKEN"),
		AccessSecret:   os.Getenv("ACCESS_SECRET"),
	}
}
