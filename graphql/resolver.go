//go:generate go run github.com/99designs/gqlgen -v

package graphql

import (
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

// func (r *Resolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
// 	return meetups, nil
// }
