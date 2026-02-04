package internals

import (
	// "errors"
	"fmt"
)

type Topic struct {
	id int64
	// producerId -> record 				// change it to the another unique key value pair
	records map[int64][]Record
	// producerId -> last record index
	lastRecordOffsets map[int64]int64
	producerIds       []int64
	consumerIds       []int64

	// consumerId -> notification chan
	consumerChans map[int64]chan RecordNotification
}

func NewTopic(id int64) *Topic {
	return &Topic{
		id:                id,
		records:           make(map[int64][]Record),
		lastRecordOffsets: map[int64]int64{},
		producerIds:       make([]int64, 0),
		consumerIds:       make([]int64, 0),
		consumerChans:     make(map[int64]chan RecordNotification),
	}
}
func (t *Topic) GetId() int64 {
	return t.id
}

func (t *Topic) Subscribe(id int64, ut UserType) chan RecordNotification {
	switch ut {
	case PRODUCER:
		t.producerIds = append(t.producerIds, id)
		return nil
	case CONSUMER:
		t.consumerIds = append(t.consumerIds, id)
		ch := make(chan RecordNotification, 10)
		t.consumerChans[id] = ch
		return ch
	}
	return nil
}

func (t *Topic) InsertRecord(producerId int64, record Record) error {
	// if !t.producerIds[producerId] {
	// 	return errors.New("you are not authenticate to insert data")
	// }

	// set record offset
	index := int64(len(t.records[producerId]))
	t.records[producerId] = append(t.records[producerId], record)

	for _, ch := range t.consumerChans {
		ch <- RecordNotification{
			TopicId:    t.id,
			ProducerId: producerId,
			Index:      index,
		}
	}

	return nil
}

func (t *Topic) RetrieveRecord(pid int64, idx int64) ([]byte, error) {
	return t.records[pid][idx].data, nil
}

func (t *Topic) PrintTopic() {

	fmt.Println("------------------------------------------")

	fmt.Printf("Topic: #%d\n", t.id)

	fmt.Printf("\tRecords: (%d)\n", len(t.records))

	for producerId, records := range t.records {
		fmt.Printf("\t\t")
		fmt.Printf("\tProducer: #%d (%d)\n\t\t\t\t", producerId, len(records))
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("\tProducers:")
	fmt.Printf("\t\t")
	for _, pid := range t.producerIds {
		fmt.Printf("%v, ", pid)
	}
	fmt.Println()

	fmt.Println("\tConsumers:")
	fmt.Printf("\t\t")
	for _, pid := range t.consumerIds {
		fmt.Printf("%v", pid)
	}
	fmt.Println()

	fmt.Println("------------------------------------------")

}
