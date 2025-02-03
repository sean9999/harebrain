package harebrain

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
)

// *dog implements EncodeHasher
var _ SERDEHasher = (*dog)(nil)

type dog struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Says string `json:"says,omitempty"`
}

type nakedCat struct {
	Id    int
	Size  int
	Breed string
}

type cat = JsonRecord[nakedCat]

func (j *dog) Hash() string {
	b := j.Serialize()
	h := crc32.ChecksumIEEE(b)
	return fmt.Sprintf("%x", h)
}
func (j *dog) Key() string {
	return j.Hash() + ".json"
}

func (j *dog) Serialize() []byte {
	return must(json.Marshal(j))
}
func (j *dog) Deserialize(p []byte) {
	err := json.Unmarshal(p, j)
	if err != nil {
		panic(err)
	}
}
func (j *dog) Clone() SERDEHasher {
	var j2 = new(dog)
	jbytes := j.Serialize()
	j2.Deserialize(jbytes)
	return j2
}

func TestTable_Insert(t *testing.T) {
	t.Run("folder that exists", func(t *testing.T) {
		db := NewDatabase()
		err := db.Open("data")
		assert.Nil(t, err, "error should be nil")
	})
	t.Run("folder that doesn't exist", func(t *testing.T) {
		db := NewDatabase()
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

	t.Run("insert some cats", func(t *testing.T) {
		db := NewDatabase()
		err := db.Open("data")
		assert.Nil(t, err, "error should be nil")
		millie := &cat{nakedCat{1, 5, "Babydoll"}}
		oliver := &cat{nakedCat{2, 9, "American House"}}
		err = db.Table("cats").Insert(millie)
		assert.Nil(t, err)
		err = db.Table("cats").Insert(oliver)
		assert.Nil(t, err)

		err = db.Table("cats").Delete(millie.Hash())
		assert.Nil(t, err)

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
		b, err := db.Table("dogs").Get("1085bb52.json")
		assert.Nil(t, err, "get record from database")
		d.Deserialize(b)
		assert.Equal(t, "Millie", d.Name, "dog's name should be Millie")
	})

}

type employee struct {
	Name   string `msgpack:"name"`
	Job    string `msgpack:"job"`
	Salary int    `msgpack:"salary"`
}

func (j *employee) Hash() string {
	return fmt.Sprintf("%s%s%d.mpack", j.Name, j.Job, j.Salary)
}
func (j *employee) Serialize() []byte {
	// type e struct {
	// 	Name   string
	// 	Job    string
	// 	Salary int
	// }
	// f := e{j.Name, j.Job, j.Salary}
	return must(msgpack.Marshal(j))
}

func (j *employee) Deserialize(p []byte) {
	err := msgpack.Unmarshal(p, j)
	if err != nil {
		panic(err)
	}
}

func TestTableCustom(t *testing.T) {

	bob := &employee{
		Name:   "bob",
		Job:    "custodian",
		Salary: 99,
	}

	db := NewDatabase()
	err := db.Open("data")
	assert.NoError(t, err)
	//table := db.Table("employees")
	err = db.Table("employees").Insert(bob)
	assert.NoError(t, err)

	b, err := db.Table("employees").Get(bob.Hash())
	assert.NoError(t, err)

	person := new(employee)

	assert.NotPanics(t, func() {
		person.Deserialize(b)
	})

	assert.NoError(t, err)

	assert.Equal(t, bob.Name, person.Name)

}
