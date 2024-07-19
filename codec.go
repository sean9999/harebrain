package harebrain

type Codec interface {
	Serialize(EncodeHasher) []byte
	Deserialize([]byte) *EncodeHasher
	Hash(EncodeHasher) string
}
