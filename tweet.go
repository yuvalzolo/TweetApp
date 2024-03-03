package main

import (
	"errors"
	"time"
)

type Tweet struct {
	Content   string
	UserName  string
	CreatedAt time.Time
}
var (
	tweets []Tweet
)

func PostTweet(userName string, tweetContent string) error {
	_, exists := users[userName]
	if !exists {
		return errors.New("user not found")
	}
	tweet := Tweet{
		Content:   tweetContent,
		UserName:  userName,
		CreatedAt: time.Now(),
	}
	tweets = append(tweets, tweet)
	return nil
}
