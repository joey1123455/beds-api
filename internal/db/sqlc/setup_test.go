package db

import (
	"database/sql"
	"os"
	"testing"

	"github.com/joey1123455/beds-api/internal/config"
	"github.com/joey1123455/beds-api/testutils"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB
var cfg config.Config

const (
	packageName = "db"
)

func TestMain(m *testing.M) {

	setup := testutils.SetupTest(packageName, "../../..", ".env.test")

	testDB = setup.TestDB
	cfg = setup.Config
	testQueries = New(testDB)

	code := m.Run()

	testutils.TeardownTest(setup, packageName)

	os.Exit(code)

}
