package node

import "io"

type Database interface {

	//reader
	Has(key []byte) (bool, error)

	Get(key []byte) ([]byte, error)

	//writer
	Put(key, value []byte) error

	Delete(key []byte) error

	Iterator

	io.Closer
}

type Iterator interface {
	Next() bool

	Error() error

	Key() []byte

	Value() []byte

	Release()
}
