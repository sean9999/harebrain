package harebrain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

// func TestDatabase_LoadTables(t *testing.T) {
// 	db := new(Database)
// 	db.Open("data")
// 	tbls, err := db.LoadTables()
// 	assert.Nil(t, err)
// 	assert.Contains(t, tbls, "cats")
// 	assert.Contains(t, tbls, "dogs")
// }
