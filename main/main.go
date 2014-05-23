package main

import (
	"fmt"
	"github.com/steaz/reddilytics/model"
	"github.com/steaz/reddilytics/update"
	"time"
)

func main() {
	for {
		err := model.SetupDB()
		if err != nil {
			panic(err)
		}

		update.UpdateAll()
		fmt.Println("\n--------------------------------------\n")

		//how often?
		time.Sleep(time.Hour * 24)
	}
}
