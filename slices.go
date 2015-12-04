package butils
import (
    "io"
    "reflect"
)

// Reads a slice of Readable objects from a reader.
func ReadSlice(r io.Reader, limit uint64, slicePtr interface{}) (err error) {

    var num uint64
    if num, err = ReadVarUint(r); err != nil { return }
    if num >= limit { return ErrLimitExceeded }
    var inum int = int(num)

    elementType := reflect.TypeOf(slicePtr).Elem().Elem()
    slice := reflect.MakeSlice(reflect.SliceOf(elementType), inum, inum)

    for i := 0; i < inum; i += 1 {
        element := slice.Index(i).Addr().Interface()
        rd, ok := element.(Readable)
        if !ok { return ErrNotReadable }
        if err = rd.Read(r); err != nil { return }
    }

    reflect.ValueOf(slicePtr).Elem().Set(slice)
    return nil
}

// Writes a slice of Writable objects into a writer.
func WriteSlice(w io.Writer, limit uint64, aSlice interface{}) (err error) {
    slice := reflect.ValueOf(aSlice)
    inum := slice.Len()

    if err = WriteVarUint(w, uint64(inum)); err != nil { return }

    for i := 0; i < inum; i += 1 {
        element := slice.Index(i).Addr().Interface()
        wr, ok := element.(Writable)
        if !ok { return ErrNotWritable }
        if err = wr.Write(w); err != nil { return }
    }
    return
}