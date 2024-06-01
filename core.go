package main

import (
	"errors"
	"sync"
)

var store = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

var ErrorNoSuchkey = errors.New("no such key")
var ErrorKeyExists = errors.New("key already exists")

func Put(key string, value string) error {
	if _, ok := store.m[key]; ok {
		return ErrorKeyExists
	}
	store.Lock()
	store.m[key] = value
	store.Unlock()
	return nil

}

func Get(key string) (string, error) {
	store.RLock()
	value, ok := store.m[key]
	if !ok {
		return " ", ErrorNoSuchkey
	}
	store.RUnlock()
	return value, nil
}

func Del(key string) error {
	if _, ok := store.m[key]; !ok {
		return ErrorNoSuchkey
	}
	delete(store.m, key)
	return nil
}
