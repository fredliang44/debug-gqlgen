package graphql

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"math/rand"
)

type Resolver struct {
}

var MessageChannel = make(chan *Message, 1)

func (r *mutationResolver) Post(ctx context.Context, text string) (*Message, error) {
	msg := Message{
		ID:   randString(10),
		Text: text,
	}
	MessageChannel <- &msg
	return &msg, nil
}

func (r *subscriptionResolver) MessageAdded(ctx context.Context) (<-chan *Message, error) {
	return MessageChannel, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
