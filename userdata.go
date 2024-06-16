package main

import (
	"github.com/ddosify/go-faker/faker"
)

var lfaker = faker.NewFaker()

var users = []*User{
	{IDField: "fe61b7d8-42f2-457c-b740-9fda21b0abc9", NameField: lfaker.RandomPersonFullName(), EmailField: lfaker.RandomEmail(), FriendIDs: []string{"1dd536af-04fc-4acb-9788-79dd15aa4f94", "7351ecc2-d5c2-4a7f-8239-126b082850fe"}},
	{IDField: "1dd536af-04fc-4acb-9788-79dd15aa4f94", NameField: lfaker.RandomPersonFullName(), EmailField: lfaker.RandomEmail(), FriendIDs: []string{}},
	{IDField: "7351ecc2-d5c2-4a7f-8239-126b082850fe", NameField: lfaker.RandomPersonFullName(), EmailField: lfaker.RandomEmail(), FriendIDs: []string{}},
}

func findUserById(id string) *User {
	for _, user := range users {
		if user.IDField == id {
			return user
		}
	}
	return nil
}

func findFriendsByUserId(userId string) []*User {
	var friends []*User
	user := findUserById(userId)
	if user == nil {
		return nil
	}

	for _, friendUserId := range user.FriendIDs {
		friend := findUserById(friendUserId)
		if friend != nil {
			friends = append(friends, friend)
		}
	}

	return friends
}
