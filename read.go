package butils
import (
    "io"
    "encoding/binary"
)

var byteOrder = binary.BigEndian

// Same as io.ReadFull(), but ignores the number of bytes read.
func ReadFull(r io.Reader, buf []byte) error {
    _, err := io.ReadFull(r, buf)
    return err
}

// Allocates a slice of size n and reads exactly as many byts into that slice.
func ReadAllocate(r io.Reader, n uint64) (data []byte, err error) {
	data = make([]byte, n)
	err = ReadFull(r, data)
	return
}

// Reads exactly one byte from the reader.
func ReadByte(r io.Reader) (byte, error) {
	data, err := ReadAllocate(r, 1)
	if err != nil { return 0, err }
    return data[0], nil
}

// Reads a single unsigned 8-bit integer from a reader.
func ReadUint8(r io.Reader) (uint8, error) {
    b, err := ReadByte(r)
    return uint8(b), err
}

// Reads a single unsigned 16-bit integer from a reader.
func ReadUint16(r io.Reader) (uint16, error) {
	data, err := ReadAllocate(r, 2)
	if err != nil { return 0, err }
	return byteOrder.Uint16(data[:]), nil
}

// Reads a single unsigned 32-bit integer from a reader.
func ReadUint32(r io.Reader) (uint32, error) {
	data, err := ReadAllocate(r, 4)
	if err != nil { return 0, err }
	return byteOrder.Uint32(data[:]), nil
}

// Reads a single unsigned 64-bit integer from a reader.
func ReadUint64(r io.Reader) (uint64, error) {
	data, err := ReadAllocate(r, 8)
	if err != nil { return 0, err }
	return byteOrder.Uint64(data[:]), nil
}

// Reads a variable-length encoded integer.
// TODO: describe the varuint precisely
func ReadVarUint(r io.Reader) (uint64, error) {
    // Integer bounds are enforced in order to ensure that each number has exactly one representation.
    v, err := ReadUint8(r)
    if err != nil { return 0, err }
    switch v {
        case 0xfd:
            v, err := ReadUint16(r)
            if err != nil { return 0, err }
            if v < 0xfd { return 0, ErrVaruintFormat }
            return uint64(v), err
        case 0xfe:
            v, err := ReadUint32(r)
            if err != nil { return 0, err }
            if v <= 0xffff { return 0, ErrVaruintFormat }
            return uint64(v), err
        case 0xff:
            v, err := ReadUint64(r)
            if err != nil { return 0, err }
            if v <= 0xffffffff { return 0, ErrVaruintFormat }
            return v, err
    }
    return uint64(v), nil
}

// Reads a slice of bytes of prefixed varuint length from the reader.
func ReadVarBytes(r io.Reader, limit uint64) ([]byte, error) {
    l, err := ReadVarUint(r);
    if err != nil { return nil, err }
    if l >= limit { return nil, ErrLimitExceeded }
    return ReadAllocate(r, l)
}

// Reads a string with prefixed varuint length from the reader.
func ReadString(r io.Reader, limit uint64) (string, error) {
    buf, err := ReadVarBytes(r, limit)
    if err != nil { return "", err }
    return string(buf), nil
}