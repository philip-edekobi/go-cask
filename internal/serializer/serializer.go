package serializer

const (
	TIMESTAMPSIZE = 8 // in bytes
	FIELDSIZE     = 4
)

type Record struct {
	CrcChecksum []byte
	Timestamp   int64
	KeySz       int
	ValueSz     int
	Key         string
	Value       string
}

type CorruptionError struct{}

func (c CorruptionError) Error() string {
	return "data does not match checksum"
}
