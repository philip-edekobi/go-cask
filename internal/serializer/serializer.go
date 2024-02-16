package serializer

const (
	TIMESTAMPSIZE = 8 // in bytes
	FIELDSIZE     = 4
	HEADERLENGTH  = 20
)

type Header struct {
	CrcChecksum []byte
	Timestamp   int64
	KeySize     int
	ValueSize   int
}

type Record struct {
	CrcChecksum []byte
	Timestamp   int64
	KeySz       int
	ValueSz     int
	Key         string
	Value       string
}

type ErrCorruptedRecord struct{}

func (c ErrCorruptedRecord) Error() string {
	return "data does not match checksum"
}

type ErrInvalidHeader struct{}

func (ih ErrInvalidHeader) Error() string {
	return "header isn't in correct format"
}
