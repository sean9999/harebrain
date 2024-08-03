package harebrain

import (
	"encoding"
	"encoding/json"
	"fmt"
	"hash/crc64"
)

// an EncodeHasher is a record in a table in a harebrain database
type EncodeHasher interface {
	Hash() string
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

var _ EncodeHasher = (*JsonRecord[string])(nil)

// JsonRecord is an EncodeHasher that serializes to JSON
type JsonRecord[T any] struct {
	Data T
}

// Hash produces random looking hex, plus a ".json" extension
func (j *JsonRecord[T]) Hash() string {
	b, _ := j.MarshalBinary()
	hash := crc64.Checksum(b, crc64.MakeTable(6996396))
	return fmt.Sprintf("%x.json", hash)
}

// MarshalBinary marshals to JSON
func (j *JsonRecord[T]) MarshalBinary() ([]byte, error) {
	return json.Marshal(j.Data)
}

// UnmarshalBinary unmarshals from JSON
func (j *JsonRecord[T]) UnmarshalBinary(p []byte) error {
	var data T
	err := json.Unmarshal(p, data)
	if err != nil {
		return err
	}
	j.Data = data
	return nil
}
