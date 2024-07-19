package harebrain

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dog struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Says string `json:"says,omitempty"`
}

func (j *dog) Hash() string {
	b, _ := j.MarshalBinary()
	h := crc32.ChecksumIEEE(b)
	return fmt.Sprintf("%x.json", h)
}
func (j *dog) MarshalBinary() ([]byte, error) {
	return json.Marshal(j)
}
func (j *dog) UnmarshalBinary(p []byte) error {
	return json.Unmarshal(p, j)
}

func TestTable_Insert(t *testing.T) {
	t.Run("folder that exists", func(t *testing.T) {
		db := new(Database)
		err := db.Open("data")
		assert.Nil(t, err, "error should be nil")
	})
	t.Run("folder that doesn't exist", func(t *testing.T) {
		db := new(Database)
		err := db.Open("Asdfasdfasdf")
		assert.NotNil(t, err, "should be an error")
	})

	t.Run("insert fido", func(t *testing.T) {

		db := NewDatabase()
		err := db.Open("data")
		assert.Nil(t, err, "error should be nil")

		fido := &dog{1, "Fido", "woof!"}

		tables, err := db.LoadTables()

		assert.Nil(t, err)

		dtable := tables["dogs"]

		fmt.Println(dtable)

		err = tables["dogs"].Insert(fido)
		assert.Nil(t, err)

	})

}
