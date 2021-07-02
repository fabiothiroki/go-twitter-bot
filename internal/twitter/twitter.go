package twitter

import (
	"log"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
)

// StatusService abstracts the twitter status api
type StatusService interface {
	Update(status string, params *twitter.StatusUpdateParams) (*twitter.Tweet, *http.Response, error)
}

// PostTweetStatusUpdate publishes the status parameter on timeline
func PostTweetStatusUpdate(statusService StatusService, status string) {
	_, resp, err := statusService.Update(status, nil)

	if err != nil {
		log.Fatal("Error publishing tweet", err)
	}

	log.Println("Post tweet response", resp.Status)
}
