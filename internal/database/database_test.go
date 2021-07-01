package database

import (
	"context"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
)

func TestGetLeastRecentPostedQuote(t *testing.T) {
	mockConn, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockConn.Close(context.Background())

	rows := pgxmock.NewRows([]string{"id", "text", "author"}).
		AddRow(1, "tweet quote", "Fabio")

	mockConn.
		ExpectQuery("select id, text, author from quotes order by last_posted_at ASC LIMIT 1").
		WillReturnRows(rows)

	dbService := &DbService{DbConn: mockConn}
	result := dbService.GetLeastRecentPostedQuote()

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, result, &Quote{ID: 1, Text: "tweet quote", Author: "Fabio"})
}

func TestUpdatePostDate(t *testing.T) {
	mockConn, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockConn.Close(context.Background())

	currentTime := time.Now()

	mockConn.
		ExpectQuery("^update quotes").
		WithArgs(currentTime, 1)

	dbService := &DbService{DbConn: mockConn}
	dbService.UpdatePostDate(1, currentTime)

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
