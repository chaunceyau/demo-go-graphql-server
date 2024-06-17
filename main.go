package main

import (
	"fmt"
	"log"
	"net/http"

	"gographqlserver/resolver"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

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
       		friends: [User]
        }
    `

	schema := graphql.MustParseSchema(schemaStr, &resolver.Resolver{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
