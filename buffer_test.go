package buffer

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Testdata struct {
	id   int
	name string
}

func Test_NewBuffer(t *testing.T) {

	capacity := (10)
	buffer := NewBuffer(capacity)
	assert.Equal(t, 0, buffer.offset, "initial length must be 0")
	assert.Equal(t, capacity, buffer.capacity, "capacity must be equal to the initial capacity argument")
}

func Test_Write(t *testing.T) {
	capacity := (10)
	buffer := NewBuffer(capacity)
	data := &Testdata{
		id:   1,
		name: "buffello",
	}

	assert.Equal(t, 0, buffer.offset, "initial length must be 0")
	buffer.Write(data)
	assert.Equal(t, 1, buffer.offset, "lenght must be 1 after writing first data")
	buffer.Write(data)
	assert.Equal(t, 2, buffer.offset, "lenght must be 2 after writing first data")

	//write more than defined capacity
	//	for i := 0; i < 20; i += 1 {
	//		buffer.Write(data)
	//	}

	//test write after buffer is full
	buffer = NewBuffer(2)
	buffer.Write(data)
	buffer.Write(data)
	//Reading all elements will result in Resetting the readAt and offset pointers
	buffer.Read()
	buffer.Read()

	assert.Equal(t, 2, buffer.readAt, "readAt should be 2")
	assert.Equal(t, 2, buffer.offset, "write offset should be 2")
	buffer.Write(data)
	assert.Equal(t, 0, buffer.readAt, "must reset readAt to 0")
	assert.Equal(t, 1, buffer.offset, "must reset offset to 0")

	buffer.Write(data)

	assert.Equal(t, 0, buffer.readAt, "read at should not change")
	assert.Equal(t, 2, buffer.Len(), "len must return 2")

	//these calls will grow the internal slice
	buffer.Write(data)

	assert.Equal(t, 3, buffer.Len(), "len must return 3")
	assert.Equal(t, 5, buffer.Capacity(), "capacity must return 3")
	buffer.Write(data)
	buffer.Write(data)
}

func Test_Read(t *testing.T) {
	capacity := (10)

	buffer := NewBuffer(capacity)

	data1 := &Testdata{
		id:   1,
		name: "buffello",
	}

	data2 := &Testdata{
		id:   2,
		name: "buffello",
	}
	data3 := &Testdata{
		id:   3,
		name: "buffello",
	}
	//write all 3 data set
	buffer.Write(data1)
	buffer.Write(data2)
	buffer.Write(data3)
	//start reading
	var data *Testdata
	var err error
	var raw interface{}

	raw, err = buffer.Read()
	assert.Nil(t, err, "error must be nil")
	data = raw.(*Testdata)
	assert.Equal(t, 1, data.id, "must return the first value inserted on first Read call")
	assert.Equal(t, 1, buffer.readAt, "must advance readAt on Read call")

	raw, err = buffer.Read()
	assert.Nil(t, err, "error must be nil")
	data = raw.(*Testdata)
	assert.Equal(t, 2, data.id, "must return the next value inserted on next Read call")
	assert.Equal(t, 2, buffer.readAt, "must advance readAt on Read call")

	raw, err = buffer.Read()
	assert.Nil(t, err, "error must be nil")
	data = raw.(*Testdata)
	assert.Equal(t, 3, data.id, "must return the next value inserted on next Read call")
	assert.Equal(t, 3, buffer.readAt, "must advance readAt on Read call")

	//further read call will reset the buffer and returons io.EOF
	raw, err = buffer.Read()
	assert.Equal(t, io.EOF, err, "error must io.EOF")
	assert.Equal(t, 0, buffer.readAt, "must not advance readAt on Read call")
	assert.Equal(t, 0, buffer.Len(), "must not advance readAt on Read call")
}

func Test_Len(t *testing.T) {
	buffer := NewBuffer(10)
	data1 := &Testdata{
		id:   1,
		name: "buffello",
	}

	data2 := &Testdata{
		id:   2,
		name: "buffello",
	}
	buffer.Write(data1)
	buffer.Write(data2)
	assert.Equal(t, 2, buffer.Len(), "must return 2 after writing 2 elements")
	buffer.Read()
	assert.Equal(t, 1, buffer.Len(), "must return 1 after reading 1 elements")
	buffer.Read()
	assert.Equal(t, 0, buffer.Len(), "must return 0 after reading all elements")
	buffer.Write(data1)
	assert.Equal(t, 1, buffer.Len(), "must return 1 after writing 1 elements")
}
