package harebrain

import (
	"fmt"
	"os"

	realfs "github.com/sean9999/go-real-fs"
)

type Database struct {
	Folder     string
	Tables     map[string]*Table
	Filesystem realfs.WritableFs
}

func NewDatabase() *Database {
	db := Database{
		Tables:     map[string]*Table{},
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
	return db.Tables[name]
}

func (db *Database) LoadTables() (map[string]*Table, error) {
	entries, err := os.ReadDir(db.Folder)
	if err != nil {
		return nil, err
	}
	tables := make(map[string]*Table, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			tables[entry.Name()] = &Table{entry.Name(), db, map[string]EncodeHasher{}}
		}
	}
	db.Tables = tables
	return tables, nil
}
