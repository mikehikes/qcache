package quickcache

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

// testing if we can create a cache store
func TestCreateCacheStore(t *testing.T) {
	qcs, err := NewQCStore()

	if err != nil {
		log.Println("Error creating cache store:", err)
		return
	}

	// does the cache store exist?
	if qcs == nil {
		t.Fatalf("Cache store is nil")
		return
	}

	fmt.Println("Created cache store")
}

func TestCreateAddValuesToCacheStore(t *testing.T) {

	log.Println("Testing the cache store")

	qcs, err := NewQCStore()

	if err != nil {
		log.Println("Error creating cache store:", err)
		return
	}

	// does the cache store exist?
	if qcs == nil {
		t.Fatalf("Cache store is nil")
		return
	}

	fmt.Println("Created cache store")

	testStringKey := "testStringKey"
	testByteKey := "testByteKey"

	testStringValue := "testStringValue"
	testByteValue := []byte("testByteValue")

	// set a string value in the cache
	err = qcs.SetStr(testStringKey, testStringValue)

	if err != nil {
		log.Println("Error setting string value in cache store:", err)
		return
	}

	// set a byte array value in the cache
	err = qcs.SetBytes(testByteKey, testByteValue)

	if err != nil {
		log.Println("Error setting byte array value in cache store:", err)
		return
	}

	fmt.Println("Added values to cache store")

	// check if the values are in the cache

	// get a string value from the cache
	strVal, err := qcs.GetStr(testStringKey)

	if err != nil {
		log.Println("Error getting string value from cache store:", err)
		return
	}

	if strVal.Type != TYPE_STRING {
		t.Fatalf("Value is not a string")
		return
	}

	strValActual := strVal.Value.(string)

	if strValActual != testStringValue {
		t.Fatalf("Value is not the expected string")
		return
	}

	// does our strVal match our testStringValue?
	if strVal.Value != testStringValue {
		t.Fatalf("String value does not match testStringValue")
		return
	}

	// get a byte array value from the cache

	byteVal, err := qcs.GetBytes(testByteKey)

	if err != nil {
		log.Println("Error getting byte array value from cache store:", err)
		return
	}

	if byteVal.Type != TYPE_BYTES {
		t.Fatalf("Value is not a byte array")
		return
	}

	byteValActual := byteVal.Value.([]byte)

	// does our byteVal match our testByteValue?

	areBytesEqual := bytes.Equal(byteValActual, testByteValue)

	if !areBytesEqual {
		t.Fatalf("Byte array value does not match testByteValue")
		return
	}

	// what are our keys in the cache?
	allKeys, err := qcs.ListAllKeys()

	if err != nil {
		log.Println("Error listing keys in cache store:", err)
		return
	}

	// is the testStringKey in the cache?
	isStringKeyInCache := false
	for _, key := range allKeys {
		if key == testStringKey {
			isStringKeyInCache = true
		}
	}

	if !isStringKeyInCache {
		t.Fatalf("String key is not in cache")
		return
	}

	// is the testByteKey in the cache?
	isByteKeyInCache := false
	for _, key := range allKeys {
		if key == testByteKey {
			isByteKeyInCache = true
		}
	}

	if !isByteKeyInCache {
		t.Fatalf("Byte key is not in cache")
		return
	}

	log.Println("Cache store is working!")

}
