package graphql

import (
	"context"
	"net/http"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/vincentmac/gqlmeetup/models"
)

const userloaderKey = "userloader"

func DataloaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userloader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*models.User, []error) {
				var users []*models.User

				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()

				if err != nil {
					return nil, []error{err}
				}

				// Fix dataloader issue: db call doesn't return the correct users in the order requested
				// see: https://www.youtube.com/watch?v=GcIVFwQv0Io&list=PLzQWIQOqeUSNwXcneWYJHUREAIucJ5UZn&index=5
				// We need to reorder the results of `user` to the order requested by `ids`
				// To do this, we'll create a map to index the users by user.ID
				u := make(map[string]*models.User, len(users))

				for _, user := range users {
					u[user.ID] = user
				}

				// Next, we'll construct a new slice containing the users
				result := make([]*models.User, len(ids))

				// and add the users in the order requested by `ids`
				for i, id := range ids {
					result[i] = u[id]
				}

				// fmt.Printf("[ids] -> %v\n[users] -> %v, %v\n", ids, users[0], users[1])
				// fmt.Printf("[result] -> %v, %v\n", result[0], result[1])

				return result, nil
			},
		}

		// add userloader onto request context
		ctx := context.WithValue(r.Context(), userloaderKey, &userloader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userloaderKey).(*UserLoader)
}
