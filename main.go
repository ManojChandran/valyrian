package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Go!"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/{key}", dataPutHandler).Methods("PUT")
	r.HandleFunc("/v1/{key}", dataGetHandler).Methods("GET")
	//	r.HandleFunc("/v1/{key}", dataDelHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

var store = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

var ErrorNoSuchkey = errors.New("no such key")
var ErrorKeyExists = errors.New("key already exists")

func dataGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value, err := Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Write([]byte(value))
}

func Get(key string) (string, error) {
	store.RLock()
	value, ok := store.m[key]
	store.RUnlock()
	if !ok {
		return "", ErrorNoSuchkey
	}
	return value, nil
}

func dataPutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Put(key string, value string) error {
	if _, ok := store[key]; ok {
		return ErrorKeyExists
	}
	store.Lock()
	store.m[key] = value
	store.Unlock()

	return nil

}

//func dataDelHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	key := vars["key"]
//	err := Del(key)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusNotFound)
//		return
//	}
//	w.WriteHeader(http.StatusNoContent)
//}
//
//func Del(key string) error {
//	if _, ok := store[key]; !ok {
//		return ErrorNoSuchkey
//	}
//	delete(store, key)
//	return nil
//}
