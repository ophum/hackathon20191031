package message

type Message struct {
	RoomId  int
	Name    string
	Message string
	Date    string
}

func NewMessage(roomId int, name, message, date string) *Message {
	return &Message{
		RoomId:  roomId,
		Name:    name,
		Message: message,
		Date:    date,
	}
}
