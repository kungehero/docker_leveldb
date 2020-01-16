package dockerleveldb

import "github.com/syndtr/goleveldb/leveldb"

//batch add data
func batchPut(db *leveldb.DB) {
	batch := new(leveldb.Batch)
	batch.Put([]byte("foo"), []byte("value"))
	batch.Put([]byte("bar"), []byte("another value"))
	batch.Delete([]byte("baz"))
	db.Write(batch, nil)
}

//batch update data
func batchUpdate(db *leveldb.DB) {
	batch := new(leveldb.Batch)
	batch.Put([]byte("foo"), []byte("value"))
	batch.Put([]byte("bar"), []byte("another value"))
	batch.Delete([]byte("baz"))
	db.Write(batch, nil)
}

//batch delete data
func batchDelete(db *leveldb.DB) {
	batch := new(leveldb.Batch)
	batch.Put([]byte("foo"), []byte("value"))
	batch.Put([]byte("bar"), []byte("another value"))
	batch.Delete([]byte("baz"))
	db.Write(batch, nil)
}
