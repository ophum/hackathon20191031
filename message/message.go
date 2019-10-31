package message

type Message struct {
	RoomId int
	Message string
	Date string
}

func NewMessage(roomId int, message, date string) *Message {
	return &Message{
		RoomId: roomId,
		Message: message,
		Date: date,
	}
}

