package harebrain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonRecord_Hash(t *testing.T) {

	fido := &dog{1, "Fido", "woof!"}
	hash := fido.Hash()

	assert.Equal(t, "fbd35d5b.json", hash, "wrong hash for dog")

}
