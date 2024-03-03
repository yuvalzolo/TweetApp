package main

import (
	"errors"
	"sort"
)

type User struct {
	Name       string
	Follows    map[string]bool
	FollowedBy map[string]bool
}

var (
	users map[string]*User
)

func init() {
	users = make(map[string]*User)
}

func CreateUser(name string) error {
	_, exists := users[name]
	if exists {
		return errors.New("user already exists")
	}
	newUser := User{
		Name:       name,
		Follows:    make(map[string]bool),
		FollowedBy: make(map[string]bool),
	}
	users[name] = &newUser

	return nil
}

func GetUser(name string) (*User, error) {
	user, exists := users[name]
	if exists {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func UpdateUser(oldName string, newName string) error {
	_, exists := users[newName]
	if exists {
		return errors.New("new user name already exists")
	}
	user, exists := users[oldName]
	if exists {
		delete(users, oldName)
		user.Name = newName
		users[newName] = user
		return nil
	}
	return errors.New("old user not found")
}

func DeleteUser(name string) error {
	_, exists := users[name]
	if !exists {
		return errors.New("user not found")
	}
	delete(users, name)
	return nil
}

func FollowUser(followerName string, followingName string) error {
	follower, existsFollower := users[followerName]
	if !existsFollower {
		return errors.New("Follower Not Found")
	}
	following, existsFollowing := users[followingName]
	if !existsFollowing {
		return errors.New("Following Not Found")
	}
	if follower.Follows[followingName] {
		return errors.New("user already follows this user")
	}
	follower.Follows[followingName] = true
	following.FollowedBy[followerName] = true
	return nil
}

func UnfollowUser(followerName string, followingName string) error {
	follower, existsFollower := users[followerName]
	following, existsFollowing := users[followingName]
	if !existsFollower || !existsFollowing {
		return errors.New("one or both users not found")
	}
	if !follower.Follows[followingName] {
		return errors.New("user is not following this user")
	}
	delete(follower.Follows, followingName)
	delete(following.FollowedBy, followerName)
	return nil
}

func GetMutualFollowers(userName1 string, userName2 string) ([]*User, error) {
	user1, error := GetUser(userName1)
	if error != nil {
		return nil, errors.New("Follower not found")
	}
	user2, error := GetUser(userName2)
	if error != nil {
		return nil, errors.New("Follower Not Found")
	}
	var mutualFollowers []*User
	for follower := range user1.FollowedBy {
		if user2.Follows[follower] {
			mutualFollower, exists := users[follower]
			if exists {
				mutualFollowers = append(mutualFollowers, mutualFollower)
			}
		}
	}
	return mutualFollowers, nil
}

func GetTopInfluencers(n int) ([]User, error) {
	if n > len(users) {
		n = len(users)
	}
	var userList []User
	for _, user := range users {
		userList = append(userList, *user)
	}
	sort.Slice(userList, func(user1_index, user2_index int) bool {
		return len(userList[user1_index].FollowedBy) > len(userList[user2_index].FollowedBy)
	})
	return userList[:n], nil
}
