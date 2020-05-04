package buffer

//IBuffer contract
type IBuffer interface {
	Slice(buffer []byte, take int) []byte
	Size() int
}

type buffer struct {
	from int
	to   int
}

//NewBuffer contructor
func NewBuffer() IBuffer {
	return &buffer{from: 0, to: 0}
}

//Slice slice a buffer peace
func (b *buffer) Slice(bfpack []byte, take int) []byte {
	b.to += take
	data := []byte{}
	if b.to <= len(bfpack) {
		data = append(data, bfpack[b.from:b.to]...)
		b.from = b.to
	}
	return data
}

//Size slice a buffer peace
func (b *buffer) Size() int {
	return b.to + 1
}
