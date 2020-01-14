package graphql

import (
	"context"

	"github.com/vincentmac/gqlmeetup/models"
)

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return r.MeetupsRepo.GetMeetups()
}

func (r *queryResolver) Users(ctx context.Context, id string) (*models.User, error) {
	return r.UsersRepo.GetUserByID(id)
}
