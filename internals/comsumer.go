package internals

import "fmt"

type Consumer struct {
	id               int64
	notificationChan chan RecordNotification
}

func NewConsumer(id int64, notificationChan chan RecordNotification) *Consumer {
	return &Consumer{
		id:               id,
		notificationChan: notificationChan,
	}
}

func (c *Consumer) Consume(topic *Topic) {
	for {
		select {
		case evt := <-c.notificationChan:
			data, _ := topic.RetrieveRecord(evt.ProducerId, evt.Index)
			fmt.Printf(
				"[Consumer %d] Topic %d | Producer %d | Data: %s\n",
				c.id, evt.TopicId, evt.ProducerId, string(data),
			)
		default:
			return
		}
	}
}

// func (c *Consumer)RetriveData()(*Record){

// }
