package main

import (
	"github.com/steaz/reddilytics/analytics"
	"github.com/steaz/reddilytics/model"
	"log"
)

func main() {
	err := model.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	analytics.Analyze()
}
