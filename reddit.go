// Package reddit implements a basic client for reddit API
package reddit

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Item represents a reddit subreddit item
type Item struct {
	Title    string
	URL      string
	Comments int `json:"num_comments"`
}

// Response represents the JSON response from the reddit API
type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

// Get retrieves the list of items from reddit subreddits
func Get(subreddit string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s.json", subreddit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Data
	}
	return items, nil
}

func (i Item) String() string {
	com := ""
	switch i.Comments {
	case 0:
		// nothing
	case 1:
		com = " (1 comment)"
	default:
		com = fmt.Sprintf(" (%d comments)", i.Comments)
	}
	return fmt.Sprintf("%s%s\n%s", i.Title, com, i.URL)
}
