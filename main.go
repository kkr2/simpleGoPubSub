package main

import (
	"pubsub/pubsub"
	"time"
)

func main() {
	//Create broker
	broker := pubsub.NewBroker()

	// Add 2 topics
	topic1 := broker.AddTopic("Topic 1")
	topic2 := broker.AddTopic("Topic 2")

	//Create 2 Subscribers per topic , than add to respective topic
	sub1t1 := pubsub.CreateNewSubscriber()
	sub2t1 := pubsub.CreateNewSubscriber()

	topic1.AddSubscriber(sub1t1)
	topic1.AddSubscriber(sub2t1)

	sub1t2 := pubsub.CreateNewSubscriber()
	sub2t2 := pubsub.CreateNewSubscriber()

	topic2.AddSubscriber(sub1t2)
	topic2.AddSubscriber(sub2t2)

	//Create msg and publish to topics

	myMessage := pubsub.NewMessage("Message 1", "Jada Jada")
	myMessage2 := pubsub.NewMessage("Message 2", "Jada Jada")

	topic1.Publish(myMessage)
	topic2.Publish(myMessage2)

	time.Sleep(time.Duration(3 * time.Second))

}
