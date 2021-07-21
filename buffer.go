package buffer

import (
	"io"
)

type Buffer struct {
	slice    []interface{}
	offset   int
	capacity int
	readAt   int
}

//NewBuffer create new buffer with given capacity of interface type
func NewBuffer(capacity int) *Buffer {
	return &Buffer{
		capacity: capacity,
		offset:   0,
		slice:    make([]interface{}, capacity),
		readAt:   0,
	}
}

//Reset reset the buffer, next write will start overwriting content currently buffer holds
func (b *Buffer) Reset() {
	b.offset = 0
	b.readAt = 0
}

func (b *Buffer) empty() bool {
	return b.offset == b.readAt
}

//Write write data at next write location.
func (b *Buffer) Write(data interface{}) {
	if b.offset == b.capacity {
		//all we need is the space to write 1 element for now
		b.Grow(1)
	}

	b.slice[b.offset] = data
	b.offset += 1
}

//Grow
func (b *Buffer) Grow(n int) {
	// Implementing using bytes.Buffer.grow()
	//https://cs.opensource.google/go/go/+/refs/tags/go1.16.6:src/bytes/buffer.go;l=117
	m := b.Len()

	if m == 0 && b.offset != 0 {
		b.Reset()
	}

	if n <= b.capacity/2-m {
		// We can slide things down instead of allocating a new
		// slice. We only need m+n <= c to slide, but
		// we instead let capacity get twice as large so we
		// don't spend all our time copying.
		copy(b.slice, b.slice[b.offset:])
	} else {
		buf := make([]interface{}, n+2*b.capacity)
		copy(buf, b.slice[b.offset:])
		b.slice = buf
		b.capacity = cap(buf)
	}
}

func (b *Buffer) remainingCapacity() int {
	return b.capacity - b.offset
}

func (b *Buffer) WriteAll(data []interface{}) {
	n := len(data)
	if n > b.remainingCapacity() {
		b.Grow(n)
	}
	copy(b.slice[b.offset:], data[:n])
	b.offset += n
}

//Slice return the underlying slice, upto offset
func (b *Buffer) Slice() []interface{} {
	return b.slice[b.readAt:b.offset]
}

//Read return value pointed by readAt pointer and advance by 1 on each call
//return io.EOF on reaching end of buffer content
func (b *Buffer) Read() (interface{}, error) {
	if b.empty() {
		//reset to 0, as we dont allow user to read any data once above condition is true
		b.Reset()
		return nil, io.EOF
	}

	data := b.slice[b.readAt]
	b.readAt += 1
	return data, nil
}

//Len return the length of buffer, number of elements filled
func (b *Buffer) Len() int {
	return b.offset - b.readAt
}

//Capacity return underlying slice capacity
func (b *Buffer) Capacity() int {
	return b.capacity
}
