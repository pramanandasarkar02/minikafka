package internals

type Producer struct {
	id int64
}

func NewProducer(producerId int64, broker *Broker) *Producer {
	return &Producer{
		id: producerId,
	}
}

func (p *Producer) InsertData(topic *Topic, data []byte) {
	records := SplitIntoRecords(data)
	for _, record := range records {
		topic.InsertRecord(p.id, record)
	}

}
