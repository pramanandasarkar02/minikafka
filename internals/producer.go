package internals

import "fmt"


type Producer struct{
	id int64
	data []byte 
}

func NewProducer(id int64) *Producer{
	return &Producer{
		id: id,
	}
}

func (p *Producer)InsertData(data []byte){
	//  split into records 

	records := SplitIntoRecords(data)

	for _, record := range records{
		fmt.Println(record.offset)
	}
	// 


}

