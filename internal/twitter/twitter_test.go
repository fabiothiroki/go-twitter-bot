package twitter

import (
	"errors"
	"net/http"
	"testing"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockStatusService struct {
	mock.Mock
}

func (m *mockStatusService) Update(status string, params *twitter.StatusUpdateParams) (*twitter.Tweet, *http.Response, error) {
	args := m.Called(status, params)

	return args.Get(0).(*twitter.Tweet), args.Get(1).(*http.Response), args.Error(2)
}

func TestPostTweetStatusUpdate(t *testing.T) {
	var params *twitter.StatusUpdateParams
	postedTweet := &twitter.Tweet{ID: 1234}

	mockService := new(mockStatusService)
	mockService.On("Update", "tweeted quote", params).Return(postedTweet, &http.Response{StatusCode: 200}, nil)

	result, _ := PostTweetStatusUpdate(mockService, "tweeted quote")

	mockService.AssertExpectations(t)
	assert.Equal(t, result, postedTweet)
}

func TestPostTweetStatusUpdateError(t *testing.T) {
	var params *twitter.StatusUpdateParams
	postedTweet := &twitter.Tweet{ID: 1234}

	mockService := new(mockStatusService)
	mockService.On("Update", "tweeted quote", params).Return(postedTweet, &http.Response{StatusCode: 200}, errors.New("error"))

	defer func() { recover() }()
	PostTweetStatusUpdate(mockService, "tweeted quote")

	t.Errorf("did not panic")
}
