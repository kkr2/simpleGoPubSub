package pubsub

import (
	"fmt"

	"github.com/google/uuid"
)

type Broker struct {
	topics map[string]*Topic // map of topic Id to topic instance
}

func NewBroker() *Broker {
	return &Broker{
		make(map[string]*Topic),
	}
}

func (b *Broker) GetTopicWithId(topicId uuid.UUID) (*Topic, error) {
	topic, ok := b.topics[topicId.String()]
	if !ok {
		return nil, fmt.Errorf("topic not found")
	}
	return topic, nil
}

func (b *Broker) AddTopic(topicName string) *Topic {
	newTopic := NewTopic(topicName)
	b.topics[newTopic.id.String()] = newTopic
	return newTopic
}
