package internals

type Topic struct {
	id      int64
	records []Record

	subscribedProducerIds []int64
	subscribedConsumerIds []int64
}

func NewTopic(id int64) *Topic {
	return &Topic{
		id:      id,
		records: make([]Record, 0),
	}
}

func (t *Topic) InsertRecord(record *Record) {
	t.records = append(t.records, *record)
}

func (t *Topic) RetriveRecord() {

}
