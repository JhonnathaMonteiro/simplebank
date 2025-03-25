package db

import (
	"database/sql"
	"jhonnatha/simplebank/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("failed to open database connection")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
