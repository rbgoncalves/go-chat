package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-chat/graph/generated"
	"go-chat/graph/model"
	"time"
)

func (r *mutationResolver) SendMessage(ctx context.Context, input model.NewMessage) (*model.ChatMessage, error) {
	msg := &model.ChatMessage{
		ID:   fmt.Sprintf("%d", time.Now().UnixNano()),
		From: input.From,
		Text: input.Text,
	}
	r.messages = append(r.messages, msg)

	return msg, nil
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.ChatMessage, error) {
	return r.messages, nil
}

func (r *subscriptionResolver) OnNewMessage(ctx context.Context) (<-chan *model.ChatMessage, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
