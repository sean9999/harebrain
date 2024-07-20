package harebrain

import (
	"encoding"
	"encoding/json"
	"fmt"
	"hash/crc64"
)

type Record[T EncodeHasher] struct {
	Data T
}

func (r *Record[T]) Hash() string {
	return r.Data.Hash()
}
func (r *Record[T]) MarshalBinary() ([]byte, error) {
	return r.Data.MarshalBinary()
}
func (r *Record[T]) UnmarshalBinary(p []byte) error {
	return r.Data.UnmarshalBinary(p)
}

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

// func (j *JsonRecord[T]) Clone() JsonRecord[T] {
// 	d1Bytes, _ := j.MarshalBinary()
// 	buf := new(bytes.Buffer)

// 	var d2 T

// 	return JsonRecord[T]{Data: j.Data}
// }
