//go:generate go run github.com/99designs/gqlgen -v

package gqlmeetup

import (
	"context"
	"errors"

	"github.com/vincentmac/gqlmeetup/models"
	"github.com/vincentmac/gqlmeetup/postgres"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

var meetups = []*models.Meetup{
	{
		ID:          "1",
		Name:        "A meetup",
		Description: "A description",
		UserId:      "1",
	},
	{
		ID:          "2",
		Name:        "A second meetup",
		Description: "A second description",
		UserId:      "2",
	},
}

var users = []*models.User{
	{
		ID:       "1",
		Username: "Bob",
		Email:    "bob@gmail.com",
	},
	{
		ID:       "2",
		Username: "Tim",
		Email:    "tim@gmail.com",
	},
}

type Resolver struct {
	MeetupsRepo postgres.MeetupsRepo
	UsersRepo   postgres.UsersRepo
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return meetups, nil
}

func (r *Resolver) Meetup() MeetupResolver {
	return &meetupResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return r.MeetupsRepo.GetMeetups()
}

type mutationResolver struct{ *Resolver }

func (m *mutationResolver) CreateMeetup(ctx context.Context, input NewMeetup) (*models.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("description not long enough")
	}

	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserId:      "1", // hardcode to bob for now
	}

	return m.MeetupsRepo.CreateMeetup(meetup)
}

type meetupResolver struct{ *Resolver }

func (m *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return m.UsersRepo.GetUserByID(obj.UserId)
}

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
