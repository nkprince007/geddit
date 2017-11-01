package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nkprince007/reddit"
)

func main() {
	subreddits := os.Args[1:]
	if len(subreddits) == 0 {
		subreddits = append(subreddits, "golang")
	}
	for _, subreddit := range subreddits {
		fmt.Println(fmt.Sprintf("subreddit: %s", subreddit))
		items, err := reddit.Get(subreddit)
		if err != nil {
			log.Fatal(err)
		}
		for _, item := range items {
			fmt.Println(item)
		}
	}
}
