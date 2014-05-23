package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

var DB *sql.DB

// execute insertion into table corresponding to subreddit in subreddits
func Save(sr *Subreddit) error {
	// each table in postgres is titled by subreddit name
	prep := fmt.Sprintf("INSERT INTO %s (id, subscribers, accounts_active, time) VALUES ($1, $2, $3, $4);", sr.Name)

	q, err := DB.Prepare(prep)

	q.Exec(sr.Name, sr.Data.Subscribers, sr.Data.Accounts_active, getTime())

	return err
}

// dials the database, returning any error
func SetupDB() error {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=password dbname=languages sslmode=disable")

	return err
}

//returns the current time in DATE-TIME format
func getTime() string {
	format := "2006-01-02 15:04:05"

	return time.Now().Format(format)
}
