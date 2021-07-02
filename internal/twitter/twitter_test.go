package twitter

import (
	"net/http"
	"testing"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/stretchr/testify/mock"
)

type mockStatusService struct {
	mock.Mock
}

func (m *mockStatusService) Update(status string, params *twitter.StatusUpdateParams) (*twitter.Tweet, *http.Response, error) {
	args := m.Called(status, params)

	return nil, &http.Response{StatusCode: 200}, args.Error(2)
}

func TestPostTweetStatusUpdate(t *testing.T) {
	var params *twitter.StatusUpdateParams

	mockService := new(mockStatusService)
	mockService.On("Update", "tweeted quote", params).Return(nil, &http.Response{StatusCode: 200}, nil)

	PostTweetStatusUpdate(mockService, "tweeted quote")

	mockService.AssertExpectations(t)
}
