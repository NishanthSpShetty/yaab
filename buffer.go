package buffer

import (
	"io"
)

type Buffer struct {
	slice    []interface{}
	length   uint64
	capacity uint64
	readAt   uint64
}

//NewBuffer create new buffer with given capacity of interface type
func NewBuffer(capacity uint64) *Buffer {
	return &Buffer{
		capacity: capacity,
		length:   0,
		slice:    make([]interface{}, capacity),
		readAt:   0,
	}
}

//Reset reset the buffer, it will overwrite any content currently buffer holds
func (b *Buffer) Reset() {
	b.length = 0
	b.readAt = 0
}

func (b *Buffer) empty() bool {
	return b.length == b.readAt
}

//Write write data at next write location.
func (b *Buffer) Write(data interface{}) {
	b.slice[b.length] = data
	b.length += 1
}

func (b *Buffer) WriteAll(data []interface{}) {
	//unimplemented
}

//Slice return the underlying slice, upto length
func (b *Buffer) Slice() []interface{} {
	return b.slice[:b.length]
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
func (b *Buffer) Len() uint64 {
	return b.length
}

//Capacity return current buffer capacity
func (b *Buffer) Capacity() uint64 {
	return b.capacity
}
