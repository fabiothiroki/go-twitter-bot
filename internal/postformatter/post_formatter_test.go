package postformatter

import (
	"testing"

	"github.com/fabiothiroki/go-twitter-bot/internal/database"
	"github.com/stretchr/testify/assert"
)

func TestFormattedMessage(t *testing.T) {
	quote := database.Quote{1, "tweet", "Fabio Hiroki"}

	result := Format(&quote)

	assert.Equal(t, result, "tweet — Fabio Hiroki #prodmgmt #productmanagement")
}
