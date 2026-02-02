package internals

var RECORD_DATA_SIZE int64 = 64 // make it 1024 at the end

type Record struct {
	data   []byte
}

func NewRecord(data []byte) *Record {
	return &Record{
		data:   data,
	}
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
		records = append(records, *NewRecord(chunk))
	}
	return records
}




type RecordNotification struct{
	TopicId int64
	ProducerId int64
	Index int64
}