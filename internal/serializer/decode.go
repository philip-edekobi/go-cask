package serializer

import (
	"encoding/binary"
)

func DecodeKV(record []byte) *Record {
	r := &Record{}
	var start, end int

	end = FIELDSIZE
	r.CrcChecksum = record[start:end]

	start = end
	end += TIMESTAMPSIZE
	r.Timestamp = int64(binary.LittleEndian.Uint64(record[start:end]))

	start = end
	end += FIELDSIZE
	r.KeySz = int(binary.LittleEndian.Uint32(record[start:end]))

	start = end
	end += FIELDSIZE
	r.ValueSz = int(binary.LittleEndian.Uint32(record[start:end]))

	start = end
	end += r.KeySz
	r.Key = string(record[start:end])

	start = end
	end += r.ValueSz
	r.Value = string(record[start:end])

	return r
}
