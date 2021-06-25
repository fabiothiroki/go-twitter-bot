module github.com/fabiothiroki/go-twitter-bot

go 1.15

// +heroku goVersion go1.15

require (
	github.com/dghubble/oauth1 v0.7.0 // indirect
	github.com/joho/godotenv v1.3.0
	internal/database v1.0.0
	internal/twitter v1.0.0

)

replace internal/twitter => ./internal/twitter

replace internal/database => ./internal/database
