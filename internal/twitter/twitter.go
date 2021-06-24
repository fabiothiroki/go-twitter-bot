package twitter

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
)

// PostTweetStatusUpdate publishes the status parameter on timeline
func PostTweetStatusUpdate(twitterClient *twitter.Client, status string) {
	_, resp, err := twitterClient.Statuses.Update(status, nil)

	if err != nil {
		log.Fatal("Error publishing tweet", err)
	}

	log.Println("Post tweet response", resp.Status)

}
