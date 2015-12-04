package butils
import (
    "bytes"
    "io"
    "reflect"
)

// Makes a Readable object read from a slice of bytes.
// Ensures that all bytes are read in the process.
func ReadAllInto(r Readable, data []byte) (err error) {
    buf := bytes.NewBuffer(data)
    if err = r.Read(buf); err != nil { return }
    if len(buf.Bytes()) > 0 { return ErrUnreadBytes }
    return
}

// Writes a writable object to a slice of bytes.
func WriteToBytes(w Writable) ([]byte, error) {
    buf := &bytes.Buffer{}
    if err := w.Write(buf); err != nil { return nil, err }
	return buf.Bytes(), nil
}

// Writes a Writable optionally into a writer, prefixing it with 0x01.
// If the passed Writable is nil, a 0x00 byte is written instead.
func WriteOptional(w io.Writer, target Writable) (err error) {
    if reflect.ValueOf(target).IsNil() { // ugh...
        return WriteByte(w, 0x00)
    } else {
        if err = WriteByte(w, 0x01); err != nil { return }
        return target.Write(w)
    }
}

// Reads a Readable from a reader, assuming it is prefixed with byte flag equal to 0x01.
// If the flag is equal to 0x00, no read is performed.
func ReadOptional(r io.Reader, target Readable) (flag byte, err error) {
    if flag, err = ReadByte(r); err != nil { return }
    if flag == 0x00 { return }
    if flag != 0x01 { return flag, ErrFlagValue }
    return flag, target.Read(r)
}

