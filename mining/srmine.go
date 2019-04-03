package mining

import (
	"encoding/json"
	"fmt"
	"github.com/walidg/reddilytics/model"
	"io/ioutil"
	"net/http"
	"time"
)

// mines the data from the reddit API and calls save
func update(sr *model.Subreddit) {
	url := "http://www.reddit.com/r/" + sr.Name + "/about.json"
	//res, err := http.Get(url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "myClient")
	res, err := client.Do(req)
	defer res.Body.Close()

	if err != nil {
		panic(err)
	}
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	
	/*
	err = ioutil.WriteFile(sr.Name + ".json", body, 0666)
	if err != nil {
		panic(err)
	}
	*/

	err = json.Unmarshal(body, &sr)
	if err == nil {
		// can do this because sr implements Stringer interface in model/subreddits.go! cool.
		//fmt.Println(sr)
	} else {
		fmt.Println("error: ", err)
	}

	err = model.Save(sr)
	if err != nil {
		fmt.Println("Error saving to database: ", err)
	}
}

// updates all subreddits in list
func UpdateAll() {

	updateList := model.GetSubreddits()

	for _, sr := range updateList {
		update(&sr)

		// need to conform to reddit API cap
		time.Sleep(time.Second * 2)
	}
}
