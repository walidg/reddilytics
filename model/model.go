package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

// represents a single row in DB
type Entry struct {
	Tablename       string
	Subscribers     int
	Accounts_active int
	Time            string
}

var DB *sql.DB

// creates the DB object, returning any error
func SetupDB() error {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=password dbname=languages sslmode=disable")

	return err
}

// execute insertion into table corresponding to subreddit in subreddits
func Save(sr *Subreddit) error {
	// each table in postgres is titled by subreddit name
	prep := fmt.Sprintf("INSERT INTO %s (id, subscribers, accounts_active, time) VALUES ($1, $2, $3, $4);", sr.Name)

	q, err := DB.Prepare(prep)
	if err != nil {
		log.Println(err)
	}

	_, err = q.Exec(sr.Name, sr.Data.Subscribers, sr.Data.Accounts_active, getTime())
	if err != nil {
		log.Println(err)
	}

	return err
}

/*
func LoadAll() (*[]Entry, error) {
	subreddits := GetSubreddits()

	for _, sr := range subreddits {

	}
}
*/

// since model knows the subreddits, better to return a slice
func LoadNewestRow(srName string) (*Entry, error) {

	stmnt := fmt.Sprintf("SELECT subscribers, accounts_active, time FROM %s ORDER BY time DESC limit 1", srName)
	row, err := prepareAndQuery(stmnt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer row.Close()

	var entry Entry
	entry.Tablename = srName
	err = row.Scan(&entry.Subscribers, &entry.Accounts_active, &entry.Time)
	if err != nil {
		log.Println("row scan error: ", err)
		return nil, err
	}
	if err = row.Err(); err != nil {
		rows.Close()
		log.Fatal(err)
	}

	return &entry, nil
}

func LoadTenNewestRows(srName string) (*[]Entry, error) {
	stmnt := fmt.Sprintf("SELECT subscribers, accounts_active, time FROM %s ORDER BY time DESC limit 10", srName)
	rows, err := prepareAndQuery(stmnt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var entries []Entry
	for rows.Next() {
		var ent Entry
		ent.Tablename = srName

		err = rows.Scan(&ent.Tablename, &ent.Subscribers, &ent.Accounts_active, &ent.Time)
		if err != nil {
			log.Println("row scan error: ", err)
			return nil, err
		}

		entries = append(entries, ent)
	}

	if err = rows.Err(); err != nil {
		rows.Close()
		log.Fatal(err)
	}

	return &entries, nil
}

func prepareAndQuery(query string) (*sql.Rows, error) {
	/*
		prepd, err := DB.Prepare(query)
		if err != nil {
			return nil, err
		}
	*/

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

//returns the current time in DATE-TIME format
func getTime() string {

	// this formatting is detailed in "time" pkg
	format := "2006-01-02 15:04:05"

	return time.Now().Format(format)
}
