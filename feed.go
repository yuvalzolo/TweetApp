package main

import (
	"errors"
	"sort"
)

func GetUserFeed(userName string) ([]Tweet, error) {
	user, exists := users[userName]
	if !exists {
		return nil, errors.New("user not found")
	}

	feed := []Tweet{}
	for _, tweet := range tweets {
		if user.Follows[tweet.UserName] {
			feed = append(feed, tweet)
		}
	}
	sort.Slice(feed, func(feed1_index, feed2_index int) bool {
		return feed[feed1_index].CreatedAt.After(feed[feed2_index].CreatedAt)
	})
	return feed, nil
}
