package graph

import (
	"go-chat/graph/model"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	messages    []*model.ChatMessage
	subscribers map[string]chan *model.ChatMessage
	mu          sync.Mutex
}
