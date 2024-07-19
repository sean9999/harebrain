package harebrain

import (
	"encoding"
	"encoding/json"
	"fmt"
	"hash/crc64"
)

type EncodeHasher interface {
	Hash() string
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type JsonEncodeHasher struct{}

func (j *JsonEncodeHasher) Hash() string {
	b, _ := j.MarshalBinary()
	tab := crc64.MakeTable(6996)
	h := crc64.Checksum(b, tab)
	return fmt.Sprintf("%x.json", h)
}
func (j *JsonEncodeHasher) MarshalBinary() ([]byte, error) {
	return json.Marshal(j)
}
func (j *JsonEncodeHasher) UnmarshalBinary(p []byte) error {
	return json.Unmarshal(p, j)
}

// type Record interface {
// 	Filename() string
// 	Table() *Table
// 	Serialize() []byte
// 	Deserialize([]byte) Record
// }

// type Record[T ReadWriteHasher] struct {
// 	table *Table[T]
// 	Data  T
// }

// func (rec Record[T]) Path() string {
// 	return filepath.Join(rec.table.Path(), rec.Data.Hash())
// }
