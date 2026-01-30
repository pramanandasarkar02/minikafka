package internals

var RECORD_DATA_SIZE int64 = 64 // make it 1024 at the end

type Record struct {
	offset int64
	data   []byte
}

func NewRecord(offset int64, data []byte) *Record {
	return &Record{
		offset: offset,
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
		records = append(records, *NewRecord(i, chunk))
	}
	return records
}
