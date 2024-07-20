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
	Clone() EncodeHasher
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
	err := json.Unmarshal(p, data)
	if err != nil {
		return err
	}
	j.Data = data
	return nil
}
func (j *JsonRecord[T]) Clone() EncodeHasher {
	var j2 = new(JsonRecord[T])
	jbytes, _ := j.MarshalBinary()
	j2.UnmarshalBinary(jbytes)
	return j2
}
