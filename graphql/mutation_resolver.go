package graphql

import (
	"context"
	"errors"
	"fmt"

	"github.com/vincentmac/gqlmeetup/models"
)

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (m *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("description not long enough")
	}

	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserId:      "2", // hardcode to bob for now
	}

	return m.MeetupsRepo.CreateMeetup(meetup)
}

func (m *mutationResolver) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetup) (*models.Meetup, error) {
	meetup, err := m.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup does not exists")
	}

	didUpdate := false

	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name is not long enough")
		}
		meetup.Name = *input.Name
		didUpdate = true
	}
	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("description is not long enough")
		}
		meetup.Description = *input.Description
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no update completed")
	}

	meetup, err = m.MeetupsRepo.Update(meetup)

	if err != nil {
		return nil, fmt.Errorf("error while updateing meetup: %v", err)
	}

	return meetup, nil
}

func (m *mutationResolver) DeleteMeetUp(ctx context.Context, id string) (bool, error) {
	meetup, err := m.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("meetup does not exists")
	}

	err = m.MeetupsRepo.Delete(meetup)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup: %v", err)
	}

	return true, nil
}
