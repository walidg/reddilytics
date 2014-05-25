package main

import (
	"fmt"
	"github.com/steaz/reddilytics/mining"
	"github.com/steaz/reddilytics/model"
	"time"
)

func main() {
	for {
		err := model.SetupDB()
		if err != nil {
			panic(err)
		}

		mining.UpdateAll()
		fmt.Println("Done: ", time.Now())
		fmt.Println("\n--------------------------------------\n")

		//how often?
		time.Sleep(time.Hour * 1)
	}
}
