package pubsub

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type Subscriber struct {
	id        uuid.UUID     // id of subscriber
	messages  chan *Message // messages channel
	topicMeta TopicMetaData // topic it is subscribed to.
	mutex     sync.RWMutex  // lock
}

type TopicMetaData struct {
	id   uuid.UUID
	name string
}

// Returns new subscriber
func CreateNewSubscriber() *Subscriber {

	newSub := &Subscriber{
		id:       uuid.New(),
		messages: make(chan *Message),
	}
	go newSub.Listen()
	return newSub
}

// Add topic info after we subscribe to a Topic
func (s *Subscriber) AddTopicMetadata(topicId uuid.UUID, topicName string) {
	s.topicMeta = TopicMetaData{id: topicId, name: topicName}
}

// Insert Message to Subscriber Channel
func (s *Subscriber) SendMessage(message *Message) {

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	s.messages <- message

}

// Creates instance that processes messages coming to subscriber channel
func (s *Subscriber) Listen() {
	for {
		if msg, ok := <-s.messages; ok {
			fmt.Printf("Subscriber %s, received message => head: %s body: %s from topic with name %s\n", s.id, msg.head, msg.body, s.topicMeta.name)
		}
	}
}
