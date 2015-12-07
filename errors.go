package butils
import "errors"

var ErrLimitExceeded = errors.New("Length limit exceeded")
var ErrVaruintFormat = errors.New("Illegal varuint format")
var ErrUnreadBytes = errors.New("Unread bytes remaining")
var ErrFlagValue = errors.New("Invalid flag value")
var ErrNotReadable = errors.New("The given interface does not implement butils.Readable")
var ErrNotWritable = errors.New("The given interface does not implement butils.Writable")
