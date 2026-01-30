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

func (r *Record)SetOffset(off int64){
	r.offset = off
}
func (r *Record)GetOffset() int64{
	return r.offset
}
func (r *Record)SetData(data []byte){
	r.data = data
}
func (r *Record)GetData() []byte {
	return  r.data
}
