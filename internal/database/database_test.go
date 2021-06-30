package database

import (
	"context"
	"testing"

	"github.com/pashagolub/pgxmock"
)

func TestCloseConnectionAfterFetchingQuote(t *testing.T) {
	mockConn, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockConn.Close(context.Background())

	mockConn.ExpectClose()

	GetLeastRecentPostedQuote(mockConn)

	if err := mockConn.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
