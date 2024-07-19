package harebrain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// type jsonCodec struct {
// 	poly uint64
// }

// func (jc jsonCodec) Serialize(rec EncodeHasher) []byte {
// 	j, _ := json.Marshal(rec)
// 	return j
// }
// func (jc jsonCodec) Deserialize(p []byte) *EncodeHasher {
// 	rec := new(EncodeHasher)
// 	json.Unmarshal(p, rec)
// 	return rec
// }
// func (jc jsonCodec) Hash(rec EncodeHasher) string {
// 	b := rec.Serialize()
// 	chk := crc64.Checksum(b, crc64.MakeTable(jc.poly))
// 	return fmt.Sprintf("%x.json", chk)
// }

func TestDatabase_Open(t *testing.T) {

	t.Run("folder that exists", func(t *testing.T) {
		db := new(Database)
		err := db.Open("data")
		assert.Nil(t, err, "error should be nil")
	})
	t.Run("folder that doesn't exist", func(t *testing.T) {
		db := new(Database)
		err := db.Open("xxdata")
		assert.NotNil(t, err, "should be an error")
	})

}

func TestDatabase_LoadTables(t *testing.T) {
	db := new(Database)
	db.Open("data")
	tbls, err := db.LoadTables()
	assert.Nil(t, err)
	noRecords := map[string]EncodeHasher{}
	want := map[string]*Table{
		"cats": &Table{Folder: "cats", Db: db, Records: noRecords},
		"dogs": &Table{Folder: "dogs", Db: db, Records: noRecords},
	}
	assert.ElementsMatch(t, tbls, want, "wanted cats and dogs")
}
