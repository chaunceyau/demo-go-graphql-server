package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// type CreateUserInput struct {
// 	name  string
// 	email string
// }

type User struct {
	IDField    string
	NameField  string
	EmailField string
	FriendIDs  []string
}

type Resolver struct{}

func (q *Resolver) CreateUser(ctx context.Context, args *struct {
	Name  string
	Email string
}) (*User, error) {
	return &User{
		IDField:    "1",
		NameField:  args.Name,
		EmailField: args.Email,
	}, nil
}

func (q *Resolver) UserById(ctx context.Context, args struct{ ID graphql.ID }) *User {
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
		input CreateUserInput {
			name: String
			email: String!
		}

        type Mutation {
			createUser(email: String!, name: String!): User
		}

        type Query {
            userById(id: ID!): User
        }

        type User {
            id: ID!
            name: String
            email: String!
            friends: [User]!
        }
    `

	schema := graphql.MustParseSchema(schemaStr, &Resolver{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
