package graphql

import (
	"context"

	"github.com/vincentmac/gqlmeetup/models"
)

type userResolver struct{ *Resolver }

func (u *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var m []*models.Meetup

	for _, meetup := range meetups {
		if meetup.UserId == obj.ID {
			m = append(m, meetup)
		}
	}

	return m, nil
}

func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}
