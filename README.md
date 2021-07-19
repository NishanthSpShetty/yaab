# buffer.go

![github action](https://github.com/NishanthSpShetty/buffer.go/actions/workflows/go.yml/badge.svg)

Implement buffer for any type. bytes.Buffer provides good functionality to play with bytes buffer. I was not able to find something similar which can hold user defined type and provide functionality such as create fixed size slice and call Write, Read, Reset on it. 

Creating a slice with `make([]Type, size)` wouldnt allow seemless insertion of values without worrying about the indexing, append would not work as it would start appending element to `size+1` location,


