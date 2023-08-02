package main

import (
	"fmt"
	"sync"
)

type KeyValueStore struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		data: make(map[string]interface{}),
	}
}

func (kv *KeyValueStore) Set(key string, value interface{}) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.data[key] = value
}

func (kv *KeyValueStore) Get(key string) (interface{}, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	value, ok := kv.data[key]
	return value, ok
}

func main() {
	kv := NewKeyValueStore()

	// Set values
	kv.Set("name", "John")
	kv.Set("age", 30)
	kv.Set("city", "New York")

	// Get values
	name, ok := kv.Get("name")
	if ok {
		fmt.Println("Name:", name)
	} else {
		fmt.Println("Name not found.")
	}

	age, ok := kv.Get("age")
	if ok {
		fmt.Println("Age:", age)
	} else {
		fmt.Println("Age not found.")
	}

	city, ok := kv.Get("city")
	if ok {
		fmt.Println("City:", city)
	} else {
		fmt.Println("City not found.")
	}

	// Trying to get a non-existing key
	address, ok := kv.Get("address")
	if ok {
		fmt.Println("Address:", address)
	} else {
		fmt.Println("Address not found.")
	}
}
