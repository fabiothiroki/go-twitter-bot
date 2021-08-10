package postformatter

import "github.com/fabiothiroki/go-twitter-bot/internal/database"

const hashTags = " #prodmgmt #productmanagement"

// Format returns the final text to be tweeted
func Format(quote *database.Quote) string {
	return quote.Text + " — " + quote.Author + hashTags
}
