package serializer

import (
	"encoding/binary"
)

func DecodeKV(record []byte) (*Record, error) {
	r := &Record{}
	var start, end int

	if !verifyHash(record) {
		return nil, ErrCorruptedRecord
	}

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

	return r, nil
}

func DecodeHeader(header []byte) (*Header, error) {
	if len(header) != 20 {
		return nil, ErrInvalidHeader{}
	}

	h := &Header{}
	var start, end int

	end = FIELDSIZE
	h.CrcChecksum = header[:end]

	start = end
	end += TIMESTAMPSIZE
	h.Timestamp = int64(binary.LittleEndian.Uint64(header[start:end]))

	start = end
	end += FIELDSIZE
	h.KeySize = int(binary.LittleEndian.Uint32(header[start:end]))

	start = end
	end += FIELDSIZE
	h.ValueSize = int(binary.LittleEndian.Uint32(header[start:end]))

	return h, nil
}
