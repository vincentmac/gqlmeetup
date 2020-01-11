package graphql

import (
	"context"

	"github.com/vincentmac/gqlmeetup/models"
)

type meetupResolver struct{ *Resolver }

func (m *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	// return m.UsersRepo.GetUserByID(obj.UserId)
	return getUserLoader(ctx).Load(obj.UserId)
}

func (r *Resolver) Meetup() MeetupResolver {
	return &meetupResolver{r}
}
