package main

import (
	"errors"
	"log"
	"net/http"
)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Go!"))
}

func main() {
	http.HandleFunc("/", helloGoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var store = make(map[string]string)
var ErrorNoSuchkey = errors.New("no such key")

func Get(key string) (string, error) {
	value, ok := store[key]
	if !ok {
		return "", ErrorNoSuchkey
	}
	return value, nil
}

func Put(key string, value string) error {
	store[key] = value
	return nil
}

func Delete(key string) error {
	if _, ok := store[key]; !ok {
		return ErrorNoSuchkey
	}
	delete(store, key)
	return nil
}
