package main

import (
	"fmt"
	"github.com/walidg/reddilytics/analytics"
	"github.com/walidg/reddilytics/model"
	"github.com/walidg/reddilytics/mining"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {

	err := model.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	mining.UpdateAll()
	fmt.Println("Done: ", time.Now())
	fmt.Println("\n--------------------------------------\n")
	
	err = model.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	preFile, err := os.Open("pretext.md")
	if err != nil {
		log.Fatal(err)
	}
	pretext, err := ioutil.ReadAll(preFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(pretext))
	analytics.Analyze()
}
