// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ChatMessage struct {
	ID   string `json:"id"`
	From string `json:"from"`
	Text string `json:"text"`
}

type NewMessage struct {
	From string `json:"from"`
	Text string `json:"text"`
}
