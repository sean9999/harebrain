package harebrain

import (
	"path/filepath"
)

type Table struct {
	Folder string
	Db     *Database
}

func (t *Table) Path() string {
	return filepath.Join(t.Db.Folder, t.Folder)
}

func (t *Table) Insert(rec SERDEHasher) error {
	b := rec.Serialize()
	fullPath := filepath.Join(t.Path(), rec.Hash())
	err := t.Db.Filesystem.WriteFile(fullPath, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (t *Table) Get(f string) ([]byte, error) {
	fullPath := filepath.Join(t.Path(), f)
	return t.Db.Filesystem.ReadFile(fullPath)
}

func (t *Table) GetAll() (map[string][]byte, error) {
	entries, err := t.Db.Filesystem.ReadDir(t.Path())
	if err != nil {
		return nil, err
	}
	m := map[string][]byte{}
	for _, entry := range entries {
		if entry.Type().IsRegular() {
			b, err := t.Db.Filesystem.ReadFile(filepath.Join(t.Path(), entry.Name()))
			if err != nil {
				return nil, err
			}
			m[entry.Name()] = b
		}
	}
	return m, nil
}

func (t *Table) Delete(f string) error {
	fullPath := filepath.Join(t.Path(), f)
	return t.Db.Filesystem.Remove(fullPath)
}
