package resolver

import (
	"context"
	"gographqlserver/model"
	"gographqlserver/service"

	"github.com/graph-gophers/graphql-go"
)

type userResolver struct {
	u *model.User
}

func (user *userResolver) Friends(ctx context.Context) *[]*userResolver {
	val := &[]*userResolver{}
	for _, friend := range service.FindFriendsByUserId(user.u.IDField) {
		*val = append(*val, &userResolver{u: friend})
	}
	return val
}

func (user *userResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(user.u.IDField)
}

func (user *userResolver) Name(ctx context.Context) *string {
	if user.u.NameField == "" {
		return nil
	}
	return &user.u.NameField
}

func (user *userResolver) Email(ctx context.Context) string {
	return user.u.EmailField
}
