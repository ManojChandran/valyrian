package main

type EventType byte

const (
	_           = iota
	EventDelete = iota
	EventPut
)

type Event struct {
	Sequence  uint64
	Eventtype EventType // Change the name of the type to avoid name conflict
	Key       string
	value     string
}

type TransactionLogger interface {
	WriteDelete(key string)
	WritePut(key, value string)
	Err() <-chan error

	LastSequence() uint64

	Run()
	Wait()
	Close() error

	ReadEvents() (<-chan Event, <-chan error)
}
