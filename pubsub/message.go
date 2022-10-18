package pubsub

type Message struct {
	head string
	body string
}

func NewMessage(head string, msg string) *Message {
	// Returns the message object
	return &Message{
		head: head,
		body: msg,
	}
}

func (m *Message) GetMessage() (head string, body string) {
	// returns the message body.
	return m.head, m.body
}
