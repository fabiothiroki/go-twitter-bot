package twitter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTwitterCreds(t *testing.T) {
	os.Setenv("CONSUMER_KEY", "key")
	os.Setenv("CONSUMER_SECRET", "consumer_secret")
	os.Setenv("ACCESS_TOKEN", "token")
	os.Setenv("ACCESS_SECRET", "access_secret")

	result := GetTwitterCreds()

	assert.Equal(t, result, &Creds{ConsumerKey: "key", ConsumerSecret: "consumer_secret", AccessToken: "token", AccessSecret: "access_secret"})
}
