package model

// structs match reddit API @ www.reddit.com/subreddit/about.json
type Subreddit struct {
	Data SubredditData
	Name string
}

type SubredditData struct {
	Display_name       string
	Url                string
	Accounts_active    int
	Subscribers        int
	Public_description string
	Description_html   string
	Header_img         string
	Header_title       string
	Title              string
	Description        string
	Name               string
}

// returns list of subreddits to watch
func GetSubreddits() []Subreddit {
	names := [...]string{
		"golang",
		// ... intentionally an array to add more later
	}

	subreddits := []Subreddit{}

	for i, _ := range names {
		subreddits = append(subreddits, newSubreddit(names[i]))
	}

	return subreddits
}

// helper returns named subreddit, required for url in update package
func newSubreddit(name string) Subreddit {
	var sr Subreddit
	sr.Name = name
	return sr
}
