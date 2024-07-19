package harebrain

import (
	"fmt"
	"path/filepath"
)

type Table struct {
	Folder  string
	Db      *Database
	Records map[string]EncodeHasher
}

func (t *Table) Path() string {
	return filepath.Join(t.Db.Folder, t.Folder)
}

func (t *Table) Insert(rec EncodeHasher) error {
	b, err := rec.MarshalBinary()
	if err != nil {
		return err
	}
	fullPath := filepath.Join(t.Path(), rec.Hash())
	err = t.Db.Filesystem.WriteFile(fullPath, b, 0644)
	if err != nil {
		return err
	}
	t.Records[rec.Hash()] = rec
	return nil
}

func (t *Table) Get(f string) (EncodeHasher, error) {
	var rec EncodeHasher
	fullPath := filepath.Join(t.Path(), f)
	p, err := t.Db.Filesystem.ReadFile(fullPath)
	if err != nil {
		return rec, err
	}
	err = rec.UnmarshalBinary(p)
	return rec, err
}

func (t *Table) Delete(f string) error {
	return t.Db.Filesystem.Remove(f)
}

func (t *Table) LoadRecords() (map[string]EncodeHasher, []error) {
	records := map[string]EncodeHasher{}
	errs := []error{}
	entries, err := t.Db.Filesystem.ReadDir(t.Path())
	if err != nil {
		return nil, []error{err}
	}
	for _, entry := range entries {
		if entry.Type().IsRegular() {
			p, err := t.Db.Filesystem.ReadFile(entry.Name())
			if err == nil {
				var rec EncodeHasher
				err := rec.UnmarshalBinary(p)
				if err != nil {
					errs = append(errs, err)
					continue
				}
				hash := rec.Hash()
				if hash != entry.Name() {
					errs = append(errs, fmt.Errorf("hash %q doesn't match filename %q", hash, entry.Name()))
				} else {
					records[hash] = rec
				}
			} else {
				errs = append(errs, err)
			}
		}
	}
	return records, errs
}
