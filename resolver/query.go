package resolver

import (
	"context"
	"gographqlserver/service"

	"github.com/graph-gophers/graphql-go"
)

func (q *Resolver) UserById(ctx context.Context, args struct{ ID graphql.ID }) *userResolver {
	return &userResolver{u: service.FindUserById(string(args.ID))}
}
