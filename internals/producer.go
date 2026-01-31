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

	// change record offset

	// save record to the producers
	for _, record := range records {
		topic.InsertRecord(p.id, record)
	}

}
