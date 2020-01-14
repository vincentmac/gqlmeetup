package graphql

import (
	"context"

	"github.com/vincentmac/gqlmeetup/models"
)

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Meetups(ctx context.Context, filter *models.MeetupFilter, limit *int, offset *int) ([]*models.Meetup, error) {
	return r.MeetupsRepo.GetMeetups(filter, limit, offset)
}

func (r *queryResolver) Users(ctx context.Context, id string) (*models.User, error) {
	return r.UsersRepo.GetUserByID(id)
}
