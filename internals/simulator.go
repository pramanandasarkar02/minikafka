package internals

import "time"

var TOPIC_ID int64 = 1000
var CONSUMER_ID int64 = 100
var PRODUCER_ID int64 = 200

func Simulation() {
	broker := NewBroker()

	PRODUCER_ID += 1
	producer1 := NewProducer(PRODUCER_ID, broker)
	PRODUCER_ID += 1
	producer2 := NewProducer(PRODUCER_ID, broker)
	PRODUCER_ID += 1
	producer3 := NewProducer(PRODUCER_ID, broker)

	TOPIC_ID += 1
	topic1 := NewTopic(TOPIC_ID)
	TOPIC_ID += 1
	topic2 := NewTopic(TOPIC_ID)

	broker.PrintBroker()

	broker.AddNewTopic(topic1)
	broker.AddNewTopic(topic2)

	broker.PrintBroker()

	topic1.Subscribe(producer1.id, PRODUCER)
	topic1.Subscribe(producer2.id, PRODUCER)
	topic2.Subscribe(producer1.id, PRODUCER)
	topic2.Subscribe(producer3.id, PRODUCER)

	CONSUMER_ID += 1
	ch1 := topic1.Subscribe(CONSUMER_ID, CONSUMER)
	consumer1 := NewConsumer(CONSUMER_ID, ch1)

	CONSUMER_ID += 1
	ch2 := topic2.Subscribe(CONSUMER_ID, CONSUMER)
	consumer2 := NewConsumer(CONSUMER_ID, ch2)

	broker.PrintBroker()

	producer1.InsertData(topic1, []byte("Jibon is a bad boy. He is playing games in mobile."))
	producer3.InsertData(topic2, []byte("Hello world"))
	producer1.InsertData(topic2, []byte("Thasin is a good Boy"))
	producer3.InsertData(topic2, []byte("Good Morning"))
	producer3.InsertData(topic2, []byte("Good Afternoon"))
	producer3.InsertData(topic2, []byte("Good Night"))


	broker.PrintBroker()

	time.Sleep(200 * time.Millisecond)

	consumer1.Consume(topic1)
	consumer2.Consume(topic2)

	broker.PrintBroker()

}
