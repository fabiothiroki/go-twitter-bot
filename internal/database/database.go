package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

// Dbconn abstracts the database connection
type Dbconn interface {
	Close(ctx context.Context) error
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
}

// OpenConnection returns a valid db connection
func OpenConnection() Dbconn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Panicf("Unable to connect to database: %v\n", err)
	}

	return conn
}

// GetLeastRecentPostedQuote returns the oldest posted quote from database
func GetLeastRecentPostedQuote(conn Dbconn) *Quote {
	defer conn.Close(context.Background())

	query := "select id, text, author from quotes order by last_posted_at ASC LIMIT 1"
	quote := &Quote{}
	conn.QueryRow(context.Background(), query).Scan(&quote.ID, &quote.Text, &quote.Author)

	log.Printf("Quote %v", quote)

	return quote
}

// UpdatePostDate updates the posted date based on quoteId
func UpdatePostDate(quoteID int) {
	conn := OpenConnection()
	defer conn.Close(context.Background())

	query := "update quotes set last_posted_at=$1 where id=$2"
	conn.QueryRow(context.Background(), query, time.Now(), quoteID)
}
