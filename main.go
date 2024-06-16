package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ddosify/go-faker/faker"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type User struct {
	IDField    string
	EmailField string
	NameField  string
	FriendIDs  []string
}

type query struct {
}

var lfaker = faker.NewFaker()

var users = []User{
	{IDField: "fe61b7d8-42f2-457c-b740-9fda21b0abc9", NameField: lfaker.RandomPersonFullName(), EmailField: lfaker.RandomEmail(), FriendIDs: []string{"1dd536af-04fc-4acb-9788-79dd15aa4f94", "7351ecc2-d5c2-4a7f-8239-126b082850fe"}},
	{IDField: "1dd536af-04fc-4acb-9788-79dd15aa4f94", NameField: lfaker.RandomPersonFullName(), EmailField: lfaker.RandomEmail(), FriendIDs: []string{}},
	{IDField: "7351ecc2-d5c2-4a7f-8239-126b082850fe", NameField: lfaker.RandomPersonFullName(), EmailField: lfaker.RandomEmail(), FriendIDs: []string{}},
}

func findUserById(id string) *User {
	for _, user := range users {
		if user.IDField == id {
			return &user
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

func (q *query) UserById(ctx context.Context, args struct{ ID graphql.ID }) *User {
	return findUserById(string(args.ID))
}

func (user *User) Friends(ctx context.Context) []*User {
	return findFriendsByUserId(user.IDField)
}

func (user *User) ID(ctx context.Context) graphql.ID {
	return graphql.ID(user.IDField)
}

func (user *User) Name(ctx context.Context) *string {
	if user.NameField == "" {
		return nil
	}
	return &user.NameField
}

func (user *User) Email(ctx context.Context) string {
	return user.EmailField
}

func main() {
	schemaStr := `
        type User {
            id: ID!
            name: String
            email: String!
            friends: [User]!
        }

        type Query {
            userById(id: ID!): User
        }
    `

	schema := graphql.MustParseSchema(schemaStr, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
