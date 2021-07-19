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

	capacity := uint64(10)
	buffer := NewBuffer(capacity)
	assert.Equal(t, uint64(0), buffer.length, "initial length must be 0")
	assert.Equal(t, capacity, buffer.capacity, "capacity must be equal to the initial capacity argument")
}

func Test_Write(t *testing.T) {
	capacity := uint64(10)
	buffer := NewBuffer(capacity)
	data := &Testdata{
		id:   1,
		name: "buffello",
	}

	assert.Equal(t, uint64(0), buffer.length, "initial length must be 0")
	buffer.Write(data)
	assert.Equal(t, uint64(1), buffer.length, "lenght must be 1 after writing first data")
	buffer.Write(data)
	assert.Equal(t, uint64(2), buffer.length, "lenght must be 2 after writing first data")
}

func Test_Read(t *testing.T) {
	capacity := uint64(10)

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
	assert.Equal(t, uint64(1), buffer.readAt, "must advance readAt on Read call")

	raw, err = buffer.Read()
	assert.Nil(t, err, "error must be nil")
	data = raw.(*Testdata)
	assert.Equal(t, 2, data.id, "must return the next value inserted on next Read call")
	assert.Equal(t, uint64(2), buffer.readAt, "must advance readAt on Read call")

	raw, err = buffer.Read()
	assert.Nil(t, err, "error must be nil")
	data = raw.(*Testdata)
	assert.Equal(t, 3, data.id, "must return the next value inserted on next Read call")
	assert.Equal(t, uint64(3), buffer.readAt, "must advance readAt on Read call")

	//further read call will reset the buffer and returons io.EOF
	raw, err = buffer.Read()
	assert.Equal(t, io.EOF, err, "error must io.EOF")
	assert.Equal(t, uint64(0), buffer.readAt, "must not advance readAt on Read call")
	assert.Equal(t, uint64(0), buffer.Len(), "must not advance readAt on Read call")

}
