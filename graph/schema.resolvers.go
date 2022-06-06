package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go-chat/graph/generated"
	"go-chat/graph/model"
	"time"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) SendMessage(ctx context.Context, input model.NewMessage) (string, error) {
	if len(input.From) == 0 {
		return "", gqlerror.Errorf("From cannot have length 0")
	}

	if len(input.Text) == 0 {
		return "", gqlerror.Errorf("Text cannot have length 0")
	}

	msg := &model.ChatMessage{
		ID:   fmt.Sprintf("%d", time.Now().UnixNano()),
		From: input.From,
		Text: input.Text,
	}
	r.messages = append(r.messages, msg)

	// Notify all active subscribers
	r.mu.Lock()

	for _, sub := range r.subscribers {
		sub <- msg
	}

	r.mu.Unlock()

	return msg.ID, nil
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.ChatMessage, error) {
	return r.messages, nil
}

func (r *subscriptionResolver) OnNewMessage(ctx context.Context, username string) (<-chan *model.ChatMessage, error) {
	msgChan := make(chan *model.ChatMessage, 1)
	id := fmt.Sprintf("%d", time.Now().UnixNano())

	if r.subscribers == nil {
		r.subscribers = make(map[string]chan *model.ChatMessage)
	}

	// This removes the subscriber once the connection terminates
	go func() {
		<-ctx.Done()
		r.mu.Lock()
		delete(r.subscribers, id)
		r.mu.Unlock()
	}()

	r.mu.Lock()
	r.subscribers[id] = msgChan
	r.mu.Unlock()

	return msgChan, nil
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
