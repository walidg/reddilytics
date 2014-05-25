package model

// structs match reddit API @ www.reddit.com/subreddit/about.json
type Subreddit struct {
	Data SubredditData
	Name string
}

type SubredditData struct {
	Accounts_active int
	Subscribers     int
	Display_name    string
	AverageActive   float32

	/*  Unused fields
	Url                string
	Public_description string
	Description_html   string
	Header_img         string
	Header_title       string
	Title              string
	Description        string
	Name               string
	*/
}

// returns list of subreddits to watch
func GetSubreddits() []Subreddit {
	names := [...]string{
		"actionscript",
		"asm",
		"c_programming",
		"cobol",
		"cpp",
		"csharp",
		"d_language",
		"delphi",
		"dotnet",
		"fortran",
		"fsharp",
		"golang",
		"haskell",
		"iosprogramming",
		"java",
		"javascript",
		"julia",
		"lisp",
		"lua",
		"matlab",
		"pascal",
		"perl",
		"php",
		"postscript",
		"python",
		"ruby",
		"scala",
		"sql",
		"visualbasic",
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
