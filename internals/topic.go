package internals

import (
	"errors"
	"fmt"
)

type Topic struct {
	id          int64
	records     map[int64]Record
	
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
		records:     make(map[int64]Record),
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

func (t *Topic) InsertRecord(producerId int64, record *Record) error {
	if !findInTheList(t.producerIds, producerId) {
		return errors.New("you are not authenticate to insert data")
	}
	t.records[record.offset] = *record
	return nil

}

func (t *Topic) RetriveRecord(offset int64, consumerId int64)([]byte, error) {
	if !findInTheList(t.consumerIds, consumerId) {
		return []byte{}, errors.New("you are not authenticate to insert data")
	}
	if record, ok := t.records[offset]; ok{
		return record.GetData(), nil
	} else{
		return []byte{}, nil
	}

}
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
	fmt.Printf("\t\t")
	for _ , rec := range t.records{
		fmt.Printf("%v, ", rec.offset)
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