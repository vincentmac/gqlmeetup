package graphql

import (
	"context"

	"github.com/vincentmac/gqlmeetup/models"
)

type userResolver struct{ *Resolver }

func (u *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	return u.MeetupsRepo.GetMeetupsForUser(obj)
}

func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}
