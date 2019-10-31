package room

import (
	"github.com/ophum/hackathon20191031/message"
)

type Room struct {
	Id int
	Name string
	Messages []message.Message
}

func NewRoom(id int, name string) *Room {
	return &Room{
		Id: id,
		Name: name,
		Messages: []message.Message{},
	}
}