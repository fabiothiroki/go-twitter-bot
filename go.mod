module github.com/fabiothiroki/go-twitter-bot

go 1.15

// +heroku goVersion go1.15

require (
	github.com/dghubble/go-twitter v0.0.0-20210609183100-2fdbf421508e // indirect
	github.com/dghubble/oauth1 v0.7.0 // indirect
	github.com/jackc/pgx/v4 v4.11.0 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/pashagolub/pgxmock v1.2.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
)

replace github.com/fabiothiroki/go-twitter-bot/internal/twitter => ./internal/twitter

replace github.com/fabiothiroki/go-twitter-bot/internal/database => ./internal/database

replace github.com/fabiothiroki/go-twitter-bot/internal/postformatter => ./internal/postformatter
