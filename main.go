package main

import (
	"fmt"
	"log"
)

func main() {
	usersToCreate := []string{"Yuval", "Bar", "Roni", "Jenny"}
	for _, userName := range usersToCreate {
		err := CreateUser(userName)
		if err != nil {
			log.Fatalf("Failed to create user %s: %v", userName, err)
		}
		fmt.Printf("User %s created successfully\n", userName)
	}
	user, err := GetUser("Yuval")
	if err != nil {
		log.Fatalf("Failed to get user Yuval: %v", err)
	}
	fmt.Printf("Got user: %s\n", user.Name)
	err = FollowUser("Yuval", "Bar")
	if err != nil {
		log.Fatalf("Failed to follow user: %v", err)
	}
	err = FollowUser("Yuval", "Jenny")
	if err != nil {
		log.Fatalf("Failed to follow user: %v", err)
	}
	err = FollowUser("Jenny", "Bar")
	if err != nil {
		log.Fatalf("Failed to follow user: %v", err)
	}
	err = FollowUser("Bar", "Yuval")
	if err != nil {
		log.Fatalf("Failed to follow user: %v", err)
	}
	err = FollowUser("Bar", "Jenny")
	if err != nil {
		log.Fatalf("Failed to follow user: %v", err)
	}
	mutualFollowers, err := GetMutualFollowers("Yuval", "Jenny")
	if err != nil {
		log.Fatalf("Failed to retrieve mutual followers: %v", err)
	}
	fmt.Println("Mutual followers of Yuval and Jenny:")
	for _, follower := range mutualFollowers {
		fmt.Println(follower.Name)
	}
	users_list, err := GetTopInfluencers(3)
	if err != nil {
		log.Fatalf("Failed to fetch followers %v", err)
	}
	fmt.Println("Most Influencers")
	for _, follower := range users_list {
		fmt.Println(follower.Name)
	}

}
