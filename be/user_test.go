package be

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var runnable = true

func init() {
	if os.Getenv("RUN_DB_TEST") != "1" {
		return
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	if len(dbHost) == 0 {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("POSTGRES_PORT")
	if len(dbPort) == 0 {
		dbPort = "5432"
	}
	dsn := fmt.Sprintf(
		"host=%s user=postgres password= dbname=ohatori port=%s sslmode=disable TimeZone=Asia/Tokyo",
		dbHost, dbPort)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	db = gormDB
}

func Test_isValidUserName(t *testing.T) {
	testcases := []struct {
		name   string
		input  string
		result bool
	}{
		{
			name:   "less characters",
			input:  "ta",
			result: false,
		},
		{
			name:   "only small",
			input:  "taro",
			result: true,
		},
		{
			name:   "only capital",
			input:  "TARO",
			result: true,
		},
		{
			name:   "use symbol1",
			input:  "ta-ro",
			result: true,
		},
		{
			name:   "use symbol2",
			input:  "ta_ro",
			result: true,
		},
		{
			name:   "use invalid symbol",
			input:  "taro!",
			result: false,
		},
		{
			name:   "use numbers",
			input:  "1234",
			result: true,
		},
		{
			name:   "use characters and numbers",
			input:  "taro1",
			result: true,
		},
		{
			name:   "use hiragana",
			input:  "たろう",
			result: false,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			result := isValidUserName(testcase.input)
			if result != testcase.result {
				t.Errorf("Unexpected result for %s: expected=%v, actual=%v\n", testcase.input, testcase.result, result)
			}
		})
	}
}

func Test_isValidPassword(t *testing.T) {
	testcases := []struct {
		name   string
		input  string
		result bool
	}{
		{
			name:   "less characters",
			input:  "ta",
			result: false,
		},
		{
			name:   "only small",
			input:  "tanakataro",
			result: false,
		},
		{
			name:   "only capital",
			input:  "TANKATARO",
			result: false,
		},
		{
			name:   "use symbol",
			input:  "TANAKA_taro1",
			result: true,
		},
		{
			name:   "only numbers",
			input:  "0123456789",
			result: false,
		},
		{
			name:   "use small and numbers",
			input:  "tanakataro1",
			result: true,
		},
		{
			name:   "use capital and numbers",
			input:  "TANAKATARO1",
			result: true,
		},
		{
			name:   "use small and capital and numbers",
			input:  "TANAKAtaro1",
			result: true,
		},
		{
			name:   "use hiragana",
			input:  "たなかたろうtaro1",
			result: true,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			result := isValidPassword(testcase.input)
			if result != testcase.result {
				t.Errorf("Unexpected result for %s: expected=%v, actual=%v\n", testcase.input, testcase.result, result)
			}
		})
	}
}

func Test_registerUser(t *testing.T) {
	if db == nil {
		t.Skip()
	}

	rand.Seed(time.Now().Unix())

	if err := db.Migrator().DropTable(&Member{}); err != nil {
		t.Fatal(err)
	}
	if err := db.Migrator().CreateTable(&Member{}); err != nil {
		t.Fatal(err)
	}

	app := &App{
		db: db,
	}

	if _, err := registerUser(app, "hoge", "duos^aev6K"); err != nil {
		t.Fatal(err)
	}
	if _, err := registerUser(app, "hoge", "duos^aev6K"); err == nil {
		t.Fatal("Duplicate user name should be error")
	}

	var members []Member
	if err := db.Find(&members).Error; err != nil {
		t.Fatal(err)
	}

	if len(members) != 1 {
		t.Errorf("User name is not unique: expectedLen=%v, actualLen=%v", 1, len(members))
	}

	if err := db.Migrator().DropTable(&Member{}); err != nil {
		t.Fatal(err)
	}
}
