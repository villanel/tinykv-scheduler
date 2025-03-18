package mvcc

import (
	"bytes"
	"encoding/binary"

	"github.com/villanel/tinykv-scheduler/kv/storage"
	"github.com/villanel/tinykv-scheduler/kv/util/codec"
	"github.com/villanel/tinykv-scheduler/kv/util/engine_util"
	"github.com/villanel/tinykv-scheduler/proto/pkg/kvrpcpb"
	"github.com/villanel/tinykv-scheduler/scheduler/pkg/tsoutil"
)

// KeyError is a wrapper type so we can implement the `error` interface.
type KeyError struct {
	kvrpcpb.KeyError
}

func (ke *KeyError) Error() string {
	return ke.String()
}

// MvccTxn groups together writes as part of a single transaction. It also provides an abstraction over low-level
// storage, lowering the concepts of timestamps, writes, and locks into plain keys and values.
type MvccTxn struct {
	StartTS uint64
	Reader  storage.StorageReader
	writes  []storage.Modify
}

func NewMvccTxn(reader storage.StorageReader, startTs uint64) *MvccTxn {
	return &MvccTxn{
		Reader:  reader,
		StartTS: startTs,
	}
}

// Writes returns all changes added to this transaction.
func (txn *MvccTxn) Writes() []storage.Modify {
	return txn.writes
}

// PutWrite records a write at key and ts.
func (txn *MvccTxn) PutWrite(key []byte, ts uint64, write *Write) {
	// Your Code Here (4A).
	put := storage.Put{
		Key:   EncodeKey(key, ts),
		Value: write.ToBytes(),
		Cf:    engine_util.CfWrite,
	}

	txn.writes = append(txn.writes, storage.Modify{Data: put})

}

// GetLock returns a lock if key is locked. It will return (nil, nil) if there is no lock on key, and (nil, err)
// if an error occurs during lookup.
func (txn *MvccTxn) GetLock(key []byte) (*Lock, error) {
	// Your Code Here (4A).
	val, err := txn.Reader.GetCF(engine_util.CfLock, key)
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, nil
	}
	lock, err := ParseLock(val)
	if err != nil {
		return nil, err
	}
	return lock, nil
}

// PutLock adds a key/lock to this transaction.
func (txn *MvccTxn) PutLock(key []byte, lock *Lock) {
	// Your Code Here (4A).
	put := storage.Put{
		Key:   key,
		Value: lock.ToBytes(),
		Cf:    engine_util.CfLock,
	}
	txn.writes = append(txn.writes, storage.Modify{Data: put})

}

// DeleteLock adds a delete lock to this transaction.
func (txn *MvccTxn) DeleteLock(key []byte) {
	// Your Code Here (4A).
	//put := new(storage.Put)
	//put.Key=key
	//put.Cf=engine_util.CfLock
	put := storage.Delete{
		Key: key,
		Cf:  engine_util.CfLock,
	}

	txn.writes = append(txn.writes, storage.Modify{Data: put})

}

// GetValue finds the value for key, valid at the start timestamp of this transaction.
// I.e., the most recent value committed before the start of this transaction.
func (txn *MvccTxn) GetValue(key []byte) ([]byte, error) {
	// Your Code Here (4A).
	iter := txn.Reader.IterCF(engine_util.CfWrite)
	defer iter.Close()
	iter.Seek(EncodeKey(key, txn.StartTS))
	if !iter.Valid() {
		return nil, nil
	}
	if !bytes.Equal(key, DecodeUserKey(iter.Item().Key())) {
		return nil, nil
	}
	val, err := iter.Item().Value()
	if err != nil {
		return nil, err
	}
	write, err := ParseWrite(val)
	if write.Kind != WriteKindPut {
		return nil, nil
	}
	return txn.Reader.GetCF(engine_util.CfDefault, EncodeKey(key, write.StartTS))
}

// PutValue adds a key/value write to this transaction.
func (txn *MvccTxn) PutValue(key []byte, value []byte) {
	// Your Code Here (4A).
	//put := new(storage.Put)
	//put.Key=EncodeKey(key,txn.StartTS)
	//put.Value=value
	//put.Cf=engine_util.CfDefault
	put := storage.Put{
		Key:   EncodeKey(key, txn.StartTS),
		Value: value,
		Cf:    engine_util.CfDefault,
	}
	txn.writes = append(txn.writes, storage.Modify{
		Data: put,
	})

}

// DeleteValue removes a key/value pair in this transaction.
func (txn *MvccTxn) DeleteValue(key []byte) {
	// Your Code Here (4A).
	//put := new(storage.Put)
	//put.Key=EncodeKey(key,txn.StartTS)
	//put.Cf=engine_util.CfDefault
	put := storage.Delete{
		Key: EncodeKey(key, txn.StartTS),
		Cf:  engine_util.CfDefault,
	}

	txn.writes = append(txn.writes, storage.Modify{Data: put})
}

// CurrentWrite searches for a write with this transaction's start timestamp. It returns a Write from the DB and that
// write's commit timestamp, or an error.
func (txn *MvccTxn) CurrentWrite(key []byte) (*Write, uint64, error) {
	// Your Code Here (4A).
	iter := txn.Reader.IterCF(engine_util.CfWrite)
	defer iter.Close()
	for iter.Seek((EncodeKey(key, TsMax))); iter.Valid(); iter.Next() {
		Item := iter.Item()
		key := Item.Key()
		if bytes.Compare(key, EncodeKey(key, txn.StartTS)) >= 0 {
			break
		}
		value, err2 := Item.Value()
		if err2 != nil {
			return nil, 0, nil
		}
		write, err2 := ParseWrite(value)
		if err2 != nil {
			return nil, 0, nil
		}
		if write.StartTS == txn.StartTS {
			return write, decodeTimestamp(key), nil
		}
	}
	return nil, 0, nil
}

// MostRecentWrite finds the most recent write with the given key. It returns a Write from the DB and that
// write's commit timestamp, or an error.
func (txn *MvccTxn) MostRecentWrite(key []byte) (*Write, uint64, error) {
	// Your Code Here (4A).
	iter := txn.Reader.IterCF(engine_util.CfWrite)
	defer iter.Close()
	iter.Seek(EncodeKey(key, TsMax))
	if !iter.Valid() {
		return nil, 0, nil
	}
	if !bytes.Equal(key, DecodeUserKey(iter.Item().Key())) {
		return nil, 0, nil
	}
	value, err := iter.Item().Value()
	if err != nil {
		return nil, 0, err
	}
	write, err := ParseWrite(value)
	if err != nil {
		return nil, 0, nil
	}
	return write, decodeTimestamp(iter.Item().Key()), nil
}

// EncodeKey encodes a user key and appends an encoded timestamp to a key. Keys and timestamps are encoded so that
// timestamped keys are sorted first by key (ascending), then by timestamp (descending). The encoding is based on
// https://github.com/facebook/mysql-5.6/wiki/MyRocks-record-format#memcomparable-format.
func EncodeKey(key []byte, ts uint64) []byte {
	encodedKey := codec.EncodeBytes(key)
	newKey := append(encodedKey, make([]byte, 8)...)
	binary.BigEndian.PutUint64(newKey[len(encodedKey):], ^ts)
	return newKey
}

// DecodeUserKey takes a key + timestamp and returns the key part.
func DecodeUserKey(key []byte) []byte {
	_, userKey, err := codec.DecodeBytes(key)
	if err != nil {
		panic(err)
	}
	return userKey
}

// decodeTimestamp takes a key + timestamp and returns the timestamp part.
func decodeTimestamp(key []byte) uint64 {
	left, _, err := codec.DecodeBytes(key)
	if err != nil {
		panic(err)
	}
	return ^binary.BigEndian.Uint64(left)
}

// PhysicalTime returns the physical time part of the timestamp.
func PhysicalTime(ts uint64) uint64 {
	return ts >> tsoutil.PhysicalShiftBits
}
