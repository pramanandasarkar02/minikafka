package internals

var TOPIC_ID int64 = 0

func Simulation() {
	broker := NewBroker()

	producer1 := NewProducer(broker)
	producer2 := NewProducer(broker)
	producer3 := NewProducer(broker)

	TOPIC_ID += 1
	topic1 := NewTopic(TOPIC_ID)
	TOPIC_ID += 1
	topic2 := NewTopic(TOPIC_ID)

	broker.PrintBroker()

	broker.AddNewTopic(topic1)
	broker.AddNewTopic(topic2)

	broker.PrintBroker()

	topic1.Subscribe(producer1.id, ProducerType)
	topic1.Subscribe(producer3.id, ConsumerType)
	topic2.Subscribe(producer1.id, ProducerType)
	topic2.Subscribe(producer2.id, ConsumerType)
	topic2.Subscribe(producer3.id, ProducerType)

	broker.PrintBroker()

	producer1.InsertData(topic1, []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."))
	producer3.InsertData(topic2, []byte("Hello world"))
	producer1.InsertData(topic2, []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."))
	producer3.InsertData(topic2, []byte("Hello world"))

	broker.PrintBroker()
	
}
