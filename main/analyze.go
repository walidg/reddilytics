package main

import (
	"fmt"
	"github.com/steaz/reddilytics/analytics"
)

func main() {
	results := analytics.Analyze()
	fmt.Println(results)
}
