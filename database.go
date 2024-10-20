package harebrain

import (
	"fmt"
	"os"

	realfs "github.com/sean9999/go-real-fs"
)

// a Database is a root folder that acts as a container for [Table]s
type Database struct {
	Folder     string
	Filesystem realfs.WritableFs
}

func NewDatabase() *Database {
	db := Database{
		Filesystem: realfs.NewWritable(),
	}
	return &db
}

func (db *Database) Open(folder string) error {
	db.Folder = folder
	info, err := os.Stat(db.Folder)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrDatabase, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("%w: path is not a folder: %q", ErrDatabase, db.Folder)
	}
	return nil
}

func (db *Database) Table(name string) *Table {
	tbl := &Table{Db: db, Folder: name}
	// info, err := os.Stat(tbl.Path())
	// if err != nil {
	// 	return nil
	// }
	// if !info.IsDir() {
	// 	return nil
	// }
	return tbl
}

// func (db *Database) LoadTables() (map[string]*Table, error) {
// 	entries, err := os.ReadDir(db.Folder)
// 	if err != nil {
// 		return nil, err
// 	}
// 	tables := make(map[string]*Table)
// 	for _, entry := range entries {
// 		if entry.IsDir() {
// 			tables[entry.Name()] = &Table{entry.Name(), db, map[string]EncodeHasher{}}
// 		}
// 	}
// 	db.Tables = tables
// 	return tables, nil
// }
