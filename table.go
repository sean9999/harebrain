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
	return nil
}

func (t *Table) Load(f string, rec EncodeHasher) error {
	fullPath := filepath.Join(t.Path(), f)
	p, err := t.Db.Filesystem.ReadFile(fullPath)
	if err != nil {
		return err
	}
	return rec.UnmarshalBinary(p)
}

func (t *Table) LoadAll(rec EncodeHasher) map[string]EncodeHasher {
	entries, err := t.Db.Filesystem.ReadDir(t.Path())
	if err != nil {
		return nil
	}
	m := map[string]EncodeHasher{}
	for _, entry := range entries {
		if entry.Type().IsRegular() {
			rc := rec.Clone()
			err = t.Load(entry.Name(), rc)
			m[entry.Name()] = rc
		}
	}
	return m
}

func (t *Table) Delete(f string) error {
	fullPath := filepath.Join(t.Path(), f)
	return t.Db.Filesystem.Remove(fullPath)
}

// func (t *Table) LoadRecords(obj EncodeHasher) []error {
// 	records := map[string]EncodeHasher{}
// 	errs := []error{}
// 	entries, err := t.Db.Filesystem.ReadDir(t.Path())
// 	if err != nil {
// 		return nil, []error{err}
// 	}
// 	for _, entry := range entries {
// 		if entry.Type().IsRegular() {
// 			p, err := t.Db.Filesystem.ReadFile(filepath.Join(t.Path(), entry.Name()))
// 			if err == nil {

// 				err := rec.UnmarshalBinary(p)
// 				if err != nil {
// 					errs = append(errs, err)
// 					continue
// 				}
// 				hash := rec.Hash()
// 				if hash != entry.Name() {
// 					errs = append(errs, fmt.Errorf("hash %q doesn't match filename %q", hash, entry.Name()))
// 				} else {
// 					records[hash] = rec
// 				}
// 			} else {
// 				errs = append(errs, err)
// 			}
// 		}
// 	}
// 	return records, errs
// }
