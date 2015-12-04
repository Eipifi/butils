package butils

import (
    "io"
    "errors"
)

// A Readable is an object that can be read (deserialized) from a stream of bytes of unknown length, represented by io.Reader.
// This object can make decisions on how many bytes to read from the buffer as the deserialization routine happens.
type Readable interface {
    Read(io.Reader) error
}

// A Writable is an object that can be written (serialized) to an io.Writer.
type Writable interface {
    Write(io.Writer) error
}

///////////////////
var ErrLimitExceeded = errors.New("Length limit exceeded")
var ErrVaruintFormat = errors.New("Illegal varuint format")
var ErrUnreadBytes = errors.New("Unread bytes remaining")
var ErrFlagValue = errors.New("Invalid flag value")
var ErrNotReadable = errors.New("The given interface does not implement butils.Readable")
var ErrNotWritable = errors.New("The given interface does not implement butils.Writable")