package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"sidecar/graph/generated"
	"sidecar/graph/model"
	"sidecar/infra/boiler"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// CreateDirectMessage is the resolver for the createDirectMessage field.
func (r *mutationResolver) CreateDirectMessage(ctx context.Context, input model.CreateDirectMessage) (*model.CreateDirectMessageResult, error) {
	return nil, nil
}

// GetDirectMessages is the resolver for the getDirectMessages field.
func (r *queryResolver) GetDirectMessages(ctx context.Context, input int) ([]*model.DirectMessagesResponse, error) {
	var DMs []*model.DirectMessagesResponse

	messages, err := boiler.DirectMessages(
		qm.Load(boiler.DirectMessageRels.User),
		qm.Where("room_id = ?", input),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	for _, m := range messages {
		dm := &model.DirectMessagesResponse{
			Element: m.Text.String,
			RoomID:  m.RoomID,
			User: &model.User{
				ID:    m.R.User.ID,
				Name:  m.R.User.Name,
				Image: &m.R.User.Image.String,
			},
		}
		DMs = append(DMs, dm)
	}

	return DMs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
