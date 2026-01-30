package internals

var PRODUCER_ID int64 = 0


type Producer struct {
	id int64
}

func NewProducer(broker *Broker) *Producer {
	PRODUCER_ID += 1
	return &Producer{
		id: PRODUCER_ID,
	}
}

func (p *Producer) InsertData(topic *Topic, data []byte) {
	//  split into records

	records := SplitIntoRecords(data)

	// // for _, record := range records{
	// // 	fmt.Println(record.offset)
	// // }

	// // send the records to the subscribed brokers
	// TOPIC_ID += 1
	// broker.AddNewTopic(p.id, TOPIC_ID, NewTopic(TOPIC_ID))

	// for _, record := range records {
	// 	broker.InsertData(TOPIC_ID, record)
	// }

	// save record to the producers
	for _, record := range records {
		topic.InsertRecord(p.id, &record)
	}

}
