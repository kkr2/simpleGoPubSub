package pubsub

import (
	"sync"

	"github.com/google/uuid"
)

type Topic struct {
	id          uuid.UUID
	name        string
	messagechan chan *Message
	subscribers []*Subscriber
	mut         sync.RWMutex // mutex lock
}

// Creates new topic instance
func NewTopic(topicName string) *Topic {
	newTopic := &Topic{
		id:          uuid.New(),
		name:        topicName,
		subscribers: []*Subscriber{},
		messagechan: make(chan *Message),
	}
	//initiate fanout Instance
	go newTopic.FanoutToSubscribers()
	return newTopic
}

// Publish message to a topic
func (t *Topic) Publish(msg *Message) {
	t.mut.RLock()
	defer t.mut.RUnlock()
	t.messagechan <- msg

}

// Attach subscriber to a topic
func (t *Topic) AddSubscriber(subscriber *Subscriber) {
	t.mut.Lock()
	defer t.mut.Unlock()
	t.subscribers = append(t.subscribers, subscriber)
	subscriber.AddTopicMetadata(t.id, t.name)
}

// Recieves a msg from topic channel,
// distributes to all subscribers
func (t *Topic) FanoutToSubscribers() {

	for {
		if msg, ok := <-t.messagechan; ok {

			if len(t.subscribers) != 0 {

				var wg sync.WaitGroup
				for _, subscriber := range t.subscribers {
					wg.Add(1)
					go func(subscriber *Subscriber) {
						defer wg.Done()
						subscriber.SendMessage(msg)

					}(subscriber)
				}
				wg.Wait()
			}

		}
	}
}
