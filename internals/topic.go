package internals

import (
	"errors"
	"fmt"
)

type Topic struct {
	id          int64
	// producerId/consumerId -> record 
	records     map[int64][]Record
	// producerId/consumerId -> last record index
	lastRecordOffsets map[int64]int64
	producerIds []int64
	consumerIds []int64
}

type UserType int

const (
	ProducerType UserType = iota
	ConsumerType
)

func NewTopic(id int64) *Topic {
	return &Topic{
		id:          id,
		records:     make(map[int64][]Record),
		lastRecordOffsets: map[int64]int64{},
		producerIds: make([]int64, 0),
		consumerIds: make([]int64, 0),
	}
}
func (t *Topic)GetId()int64{
	return t.id
}


func (t *Topic) Subscribe(id int64, ut UserType) {
	switch ut {
	case ProducerType:
		t.producerIds = append(t.producerIds, id)
	case ConsumerType:
		t.consumerIds = append(t.consumerIds, id)
	}

}

func (t *Topic) Unsubscribe(id int64, ut UserType) {
	switch ut {
	case ProducerType:
		removeFromList(t.producerIds, id)
	case ConsumerType:
		removeFromList(t.consumerIds, id)
	}

}

func (t *Topic) InsertRecord(producerId int64, record Record) error {
	if !findInTheList(t.producerIds, producerId) {
		return errors.New("you are not authenticate to insert data")
	}

	// set record offset 
	
	record.SetOffset(t.lastRecordOffsets[producerId])

	t.records[producerId] = append(t.records[producerId], record)
	t.lastRecordOffsets[producerId] += RECORD_DATA_SIZE
	return nil
}

// func (t *Topic) RetriveRecord(offset int64, consumerId int64)([]byte, error) {
// 	if !findInTheList(t.consumerIds, consumerId) {
// 		return []byte{}, errors.New("you are not authenticate to insert data")
// 	}
// 	if record, ok := t.records[consumerId]; ok{
// 		return record., nil
// 	} else{
// 		return []byte{}, nil
// 	}

// }
func removeFromList(list []int64, val int64) []int64 {
	var res []int64
	for _, value := range list {
		if val != value {
			res = append(res, value)
		}
	}
	return res
}

func findInTheList(list []int64, val int64) bool {
	for _, item := range list {
		if item == val {
			return true
		}
	}
	return false
}


func (t *Topic)PrintTopic(){

	fmt.Println("------------------------------------------")

	fmt.Printf("Topic: #%d\n", t.id)
	
	fmt.Printf("\tRecords: (%d)\n", len(t.records))
	
	for producerId , records := range t.records{
		fmt.Printf("\t\t")
		fmt.Printf("\tProducer: #%d (%d)\n\t\t\t\t", producerId, len(records))

		for _, rec := range records{
			fmt.Printf("%v, ", rec.GetOffset())
		}
		fmt.Println()
	}
	fmt.Println()
	
	
	fmt.Println("\tProducers:")
	fmt.Printf("\t\t")
	for _, pid := range t.producerIds{
		fmt.Printf("%v, ", pid)
	}
	fmt.Println()

	fmt.Println("\tConsumers:")
	fmt.Printf("\t\t")
	for _, pid := range t.consumerIds{
		fmt.Printf("%v", pid)
	}
	fmt.Println()

	fmt.Println("------------------------------------------")

}




func SplitIntoRecords(data []byte) []Record {

	records := make([]Record, 0)
	totalDataLen := int64(len(data))

	for i := int64(0); i < totalDataLen; i += RECORD_DATA_SIZE {

		end := i + RECORD_DATA_SIZE
		if end > totalDataLen {
			end = totalDataLen
		}
		chunk := data[i:end]
		records = append(records, *NewRecord(i, chunk))
	}
	return records
}
