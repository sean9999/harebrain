package harebrain

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// *dog implements EncoderHasher
var _ EncodeHasher = (*dog)(nil)

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

	t.Run("insert some dogs", func(t *testing.T) {
		db := NewDatabase()
		err := db.Open("data")
		assert.Nil(t, err, "error should be nil")
		fido := &dog{1, "Fido", "woof!"}
		dog2 := &dog{2, "Millie", "rowoo!"}
		dog3 := &dog{3, "Charles", "ratcha!"}
		err = db.Table("dogs").Insert(fido)
		assert.Nil(t, err)
		db.Table("dogs").Insert(dog2)
		db.Table("dogs").Insert(dog3)
	})

	t.Run("a file for fido exists", func(t *testing.T) {
		info, err := os.Stat("data/dogs/fbd35d5b.json")
		assert.Nil(t, err)
		assert.True(t, info.Mode().IsRegular())
	})

	t.Run("remove fido", func(t *testing.T) {
		db := NewDatabase()
		err := db.Open("data")
		assert.Nil(t, err)
		//tables, err := db.LoadTables()
		//assert.Nil(t, err)
		err = db.Table("dogs").Delete("fbd35d5b.json")
		assert.Nil(t, err)
	})

	t.Run("fido is gone", func(t *testing.T) {
		info, err := os.Stat("data/dogs/fbd35d5b.json")
		assert.NotNil(t, err)
		assert.Nil(t, info)
	})

	t.Run("Millie remains", func(t *testing.T) {
		db := NewDatabase()
		err := db.Open("data")
		assert.Nil(t, err)
		var d dog
		err = db.Table("dogs").Load("1085bb52.json", &d)
		assert.Nil(t, err)
		assert.Equal(t, "Millie", d.Name, "dog's name should be Millie")
	})

}
