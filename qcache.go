package quickcache

import (
	"fmt"
	"time"
)

// A small utility for an in-memory cache
// TODO: use generics instead of interface{}

// constants for the type of value stored in the cache
const (
	TYPE_STRING int = 0 // value is a string
	TYPE_BYTES  int = 1 // value is a byte array
)

// structure that describes a value in the cache
type QCacheVal struct {
	Type      int         // type of value stored in the cache (string or byte array)
	Value     interface{} // value stored in the cache
	Timestamp int64       // timestamp when the value was added
}

// structure that describes the method required of a cache store
type IQCacheStore interface {
	ListKeys() ([]string, error)             // list all keys in the cache
	ListStrKeys() ([]string, error)          // list all keys in the cache that are strings
	ListBytesKeys() ([]string, error)        // list all keys in the cache that are byte arrays
	GetStr(key string) (string, error)       // get a string value from the cache
	SetStr(key string, value string) error   // set a string value in the cache
	GetBytes(key string) ([]byte, error)     // get a byte array value from the cache
	SetBytes(key string, value []byte) error // set a byte array value in the cache
	HasKey(key string) (bool, int64, error)  // check if a key exists in the cache
	DeleteKey(key string) error              // delete a key from the cache
}

//  The quickcache structure store
type QCacheStore struct {
	data map[string]QCacheVal // map of key to value
	// future struct fields go here
}

// create a new cache store
func NewQCStore() (*QCacheStore, error) {
	data := make(map[string]QCacheVal)
	qcs := QCacheStore{data: data}
	return &qcs, nil
}

// list all keys in the cache (keys are strings)
func (qcs *QCacheStore) ListAllKeys() ([]string, error) {
	keys := make([]string, 0)
	for k := range qcs.data {
		keys = append(keys, k)
	}
	return keys, nil
}

// list all keys in the cache that are strings
func (qcs *QCacheStore) ListStrKeys() ([]string, error) {
	keys := make([]string, 0)
	for k := range qcs.data {
		if qcs.data[k].Type == TYPE_STRING {
			keys = append(keys, k)
		}
	}
	return keys, nil
}

// list all keys in the cache that are byte arrays
func (qcs *QCacheStore) ListBytesKeys() ([]string, error) {
	keys := make([]string, 0)
	for k := range qcs.data {
		if qcs.data[k].Type == TYPE_BYTES {
			keys = append(keys, k)
		}
	}
	return keys, nil
}

// check if a key exists in the cache
func (qcs *QCacheStore) HasKey(key string) (bool, int64, error) {
	val, ok := qcs.data[key]
	if !ok {
		return false, 0, nil
	}
	return true, val.Timestamp, nil
}

// get a key that represents a string value from the cache
func (qcs *QCacheStore) GetStr(key string) (*QCacheVal, error) {
	val, ok := qcs.data[key]
	if !ok {
		return nil, fmt.Errorf("key not found")
	}
	if val.Type != TYPE_STRING {
		return nil, fmt.Errorf("value is not a string")
	}
	return &val, nil
}

// set a key that represents a string value in the cache
func (qcs *QCacheStore) SetStr(key string, value string) error {
	currentUnixEpoch := int64(time.Now().Unix())
	qcs.data[key] = QCacheVal{TYPE_STRING, value, currentUnixEpoch}
	return nil
}

// get a key that represents a byte array value from the cache
func (qcs *QCacheStore) GetBytes(key string) (*QCacheVal, error) {
	val, ok := qcs.data[key]
	if !ok {
		return nil, fmt.Errorf("key not found")
	}
	if val.Type != TYPE_BYTES {
		return nil, fmt.Errorf("value is not a byte array")
	}
	return &val, nil
}

// set a key that represents a byte array value in the cache
func (qcs *QCacheStore) SetBytes(key string, value []byte) error {
	currentUnixEpoch := int64(time.Now().Unix())
	qcs.data[key] = QCacheVal{TYPE_BYTES, value, currentUnixEpoch}
	return nil
}

func (qcs *QCacheStore) DeleteKey(key string) error {
	delete(qcs.data, key)
	return nil
}
