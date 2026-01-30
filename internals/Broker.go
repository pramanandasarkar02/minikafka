package internals

import "fmt"

var BROKER_ID int64 = 0

type Broker struct {
	id int64
	// topicID -> prducerIds/consumerIds
	topics map[int64]*Topic
}

func NewBroker() *Broker {
	BROKER_ID += 1
	return &Broker{
		id:     BROKER_ID,
		topics: make(map[int64]*Topic),
	}
}

func (b *Broker) AddNewTopic(topic *Topic) {
	b.topics[topic.id] = topic
}


func (b *Broker) PrintBroker() {
	fmt.Println("==========================================")
	fmt.Printf("Broker details: #%d\n", b.id)
	for _, topic := range b.topics {
		topic.PrintTopic()
	}

	fmt.Println("==========================================")

	
}
