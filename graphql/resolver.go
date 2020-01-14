//go:generate go run github.com/99designs/gqlgen -v

package graphql

import (
	"github.com/vincentmac/gqlmeetup/postgres"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	MeetupsRepo postgres.MeetupsRepo
	UsersRepo   postgres.UsersRepo
}
