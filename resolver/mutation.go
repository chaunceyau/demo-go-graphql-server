package resolver

import (
	"context"
	"gographqlserver/model"
)

func (q *Resolver) CreateUser(ctx context.Context, args *struct {
	Name  string
	Email string
}) (*userResolver, error) {
	return &userResolver{
		u: &model.User{
			IDField:    "1",
			NameField:  args.Name,
			EmailField: args.Email,
		},
	}, nil
}
