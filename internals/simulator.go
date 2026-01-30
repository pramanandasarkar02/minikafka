package internals

func Simulation() {
	broker := NewBroker()

	producer := NewProducer(broker)

	broker.PrintBroker()

	producer.InsertData(broker, []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."))

	broker.PrintBroker()
}
