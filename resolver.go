//go:generate go run github.com/99designs/gqlgen -v

package gqlmeetup

import (
	"context"
	"errors"

	"github.com/vincentmac/gqlmeetup/models"
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

type Resolver struct{}

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
	return meetups, nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateMeetup(ctx context.Context, input NewMeetup) (*models.Meetup, error) {
	panic("not implemented")
}

type meetupResolver struct{ *Resolver }

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	user := new(models.User)

	for _, u := range users {
		if u.ID == obj.UserId {
			user = u
			break
		}
	}

	if user == nil {
		return nil, errors.New("user with id not found")
	}
	return user, nil
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
