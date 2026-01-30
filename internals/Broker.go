package internals

import "fmt"

var BROKER_ID int64 = 0

type Broker struct {
	id int64
	// topicID -> prducerIds/consumerIds
	subscribedProducerIds map[int64][]int64
	subscribedConsumerIds map[int64][]int64

	topics map[int64]*Topic
	// vulnerable if so frequent topics is created use lock later
}

func NewBroker() *Broker {
	BROKER_ID += 1
	return &Broker{
		id:                    BROKER_ID,
		subscribedProducerIds: make(map[int64][]int64),
		subscribedConsumerIds: make(map[int64][]int64),
		topics:                make(map[int64]*Topic),
	}
}

func (b *Broker) AddNewTopic(producerId, topicId int64, topic *Topic) {
	b.topics[topicId] = topic
	b.subscribedProducerIds[topicId] = make([]int64, 0)
	b.subscribedProducerIds[topicId] = append(b.subscribedProducerIds[topicId], producerId)
}

func (b *Broker) AddProducer(topicId, producerId int64) {
	b.subscribedProducerIds[topicId] = append(b.subscribedProducerIds[topicId], producerId)
}

func (b *Broker) InsertData(topicId int64, record Record) {
	b.topics[topicId].InsertRecord(&record)
	// notify consumer
}

func (b *Broker) PrintBroker() {
	fmt.Printf("Broker details: #%d\n", b.id)
	fmt.Println("\tTopics: ")
	for topicId, topic := range b.topics {
		fmt.Printf("\t\ttopic #%d\n\t\ttopic len: %v\n", topicId, len(topic.records))
	}

	fmt.Println("\tProducerIds:")
	for _, pid := range b.subscribedProducerIds {
		fmt.Printf("\t\tproducer: #%d\n", pid)
	}
	fmt.Println("\tConsumerIds:")
	for _, pid := range b.subscribedConsumerIds {
		fmt.Printf("\t\tconsumer: #%d\n", pid)
	}
}
