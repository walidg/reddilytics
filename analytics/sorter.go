package analytics

import (
	"fmt"
	"github.com/reddilytics/model"
)

// load most recent data.

func Analyze() string {
	return ""
}

func sortBySubs(*[]model.Subreddit) {

}

func sortByActivity(*[]model.Subreddit) {
	// this is more complicated, because I need the average
	// activity over time (past 10 days?), not a simple most recent entry pull

}

func calcActivity(sr model.Subreddit) {
	return averageActivity(sr model.Subreddit) / sr.Data.Subscribers
}

func averageActivity(sr model.Subreddit) {
	// has to make a pull on the 10? most recent database entries
}
