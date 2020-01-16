package dockerleveldb

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type LevelKv struct {
	key   string
	value string
}

//Iterate over database content:
func Iterate(db *leveldb.DB) ([]*LevelKv, error) {
	var kvs []*LevelKv
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		kvs = append(kvs, &LevelKv{key: string(k), value: string(v)})
	}
	iter.Release()

	if err := iter.Error(); err != nil {
		return nil, err
	}
	return kvs, nil
}

//Seek-then-Iterate:
func seekThenIterate(db *leveldb.DB, key []byte) (map[string][]string, error) {
	mkv := make(map[string][]string)
	iter := db.NewIterator(nil, nil)
	for ok := iter.Seek(key); ok; ok = iter.Next() {
		v := iter.Value()
		mkv[string(key)] = append(mkv[string(key)], string(v))
	}
	iter.Release()
	if err := iter.Error(); err != nil {
		return nil, err
	}
	return mkv, nil
}

//Iterate over subset of database content:
func iterateOverSubset(db *leveldb.DB) {
	iter := db.NewIterator(&util.Range{Start: []byte("foo"), Limit: []byte("xoo")}, nil)
	for iter.Next() {
		// Use key/value.
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Println(err)
	}
}

//Iterate over subset of database content with a particular prefix:
func iterateOverSubsetWithPrefix(db *leveldb.DB) {
	iter := db.NewIterator(util.BytesPrefix([]byte("foo-")), nil)
	for iter.Next() {
		// Use key/value.
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Println(err)
	}
}
