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

type JsonRecord[T any] struct {
	Data T
}

func (j *JsonRecord[T]) Hash() string {
	b, _ := j.MarshalBinary()
	tab := crc64.MakeTable(6996)
	h := crc64.Checksum(b, tab)
	return fmt.Sprintf("%x.json", h)
}
func (j *JsonRecord[T]) MarshalBinary() ([]byte, error) {
	return json.Marshal(j.Data)
}
func (j *JsonRecord[T]) UnmarshalBinary(p []byte) error {
	var data T
	return json.Unmarshal(p, data)
}
