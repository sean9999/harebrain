package harebrain

import (
	"encoding/json"
	"fmt"
	"hash/crc64"
)

// a SERDE is a Serializer and Deserializer
type SERDE interface {
	Serialize() []byte
	Deserialize([]byte)
}

// a Hasher can produce a unique deterministic string representation of itself
type Hasher interface {
	Hash() string
}

// an SERDEHasher is a record in a table in a harebrain database
type SERDEHasher interface {
	Hasher
	SERDE
}

var _ SERDEHasher = (*JsonRecord[string])(nil)

// JsonRecord is an EncodeHasher that serializes to JSON
type JsonRecord[T any] struct {
	Data T
}

func must[T any](a T, e error) T {
	if e != nil {
		panic(e)
	}
	return a
}

// Hash produces random looking hex
func (j *JsonRecord[T]) Hash() string {
	b := j.Serialize()
	hash := crc64.Checksum(b, crc64.MakeTable(6996396))
	return fmt.Sprintf("%x.json", hash)
}

// Serialize marshals to JSON, and panics on error
func (j *JsonRecord[T]) Serialize() []byte {
	return must(json.Marshal(j.Data))
}

// Deserialize unmarshals from JSON, and panics on error
func (j *JsonRecord[T]) Deserialize(p []byte) {
	var data T
	err := json.Unmarshal(p, data)
	if err != nil {
		panic(err)
	}
	j.Data = data
}
