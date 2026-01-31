package internals

import "fmt"



type Consumer struct {
	id int64
	notificationChan chan RecordNotification
}

func NewConsumer(id int64, notificationChan chan RecordNotification)*Consumer{
	return &Consumer{
		id: id,
		notificationChan: notificationChan,
	}
}

func (c *Consumer)retriveRecord(topic *Topic, producerId int64, offset int64)([]byte, error){
	return topic.RetriveRecord(c.id, producerId, offset)
}



func (c *Consumer)RetriveData(topic *Topic)[]byte{
	for event := range c.notificationChan {
		if topic.id == event.TopicId{
			data, _ := c.retriveRecord(topic, event.ProducerId, event.Offset)
			
			fmt.Printf("Topic #%d: Get data from %d offset %d\n\tdata: %v", topic.id, event.ProducerId, event.Offset, data)
		}
		
	}
	return []byte{}
}





// func (c *Consumer)RetriveData()(*Record){

// }
