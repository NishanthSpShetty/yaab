# yaab- Yet Another Anytype Buffer

![github action](https://github.com/NishanthSpShetty/yaab/actions/workflows/go.yml/badge.svg)

## Motivation
Creating a slice/array of fixed size became too painful to manage as I had to manage index manually, bytes.Buffer provided a way to allocate fixed size buffer and call Write(), I dint have to worry about indexing, filling up the buffer and managing resizing it. 
I wanted something similar to user defined types, where I can allocate predefined memory and start writing to it, while the library managed indexing, resizing for me. 

Creating a slice with `make([]Type, size)` and calling append would not work as it would start appending element to `size+1` location,

## What it is?
The intention was to build something similar to [bytes.Buffer](https://pkg.go.dev/bytes#Buffer), I have limited to exposing few API's I wanted.
key difference apart from it holds user defined data types, 
1. it doesnt take slice from users, you can only ask it to create a fixed size buffer for you, where as bytes.Buffer can be created from existing `[]byte`
2. User can write any type of data, may be generics would help us to achive writing same type of data.

## Usage

### Import

```
go get github.com/nishanthspshetty/yaab@v0.2.0
```

### Creating new buffer, writing and reading data

```
//create new buffer with size 10
buf := buffer.NewBuffer(10)
//start writing data to buffer
type SomeStruct struct {
	Name string
	Id   int
}

data := SomeStruct{
	Name: "nishanth",
	Id:   3243,
}
buf.Write(data)

//Read back the data, it returns the element of type interface and error
raw, err := buf.Read()
//when we reach end of readable content in buffer, Read will return io.EOF
if err != io.EOF {
	data = raw.(SomeStruct)
}
```

### buffer stat
```
//get the capacity of underlying slice used by the buffer
buffer.Cap() 

//number of current unread data in buffer
buffer.Len()
```

### License
GNU General Public License v3.0, [refer](https://github.com/NishanthSpShetty/yaab/blob/main/LICENSE.md)
