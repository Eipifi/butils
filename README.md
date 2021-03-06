# Butils

Butils is a handy library for handling `io.Reader` and `io.Writer`.
It introduces the following interfaces:

```go
// A Readable is an object that can be read (deserialized) from a stream of bytes of unknown length, represented by io.Reader.
// This object can make decisions on how many bytes to read from the buffer as the deserialization routine happens.
type Readable interface {
    Read(io.Reader) error
}

// A Writable is an object that can be written (serialized) to an io.Writer.
type Writable interface {
    Write(io.Writer) error
}
```

Documentation: https://godoc.org/github.com/Eipifi/butils
