package maps

import (
	"fmt"
)

// in dictionaries, keys must be of comparable types
// bool, int, float, complex, string, pointer, channel -> comparable
// struct -> comparable (all fields are comparable)
// array -> comparable (array elements are comparable)

// maps can have nil value, and they read like empty maps
// write attempt to nil value maps will give error

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound  = DictionaryErr("key not found")
	ErrKeyExists = DictionaryErr("key already exists")
)

func (dictionary Dictionary) Search(word string) (string, error) {
	// map in go returns a second value of type boolean
	value, found := dictionary[word]
	fmt.Print(value)

	if !found {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrKeyExists
	default:
		return err
	}

	return nil
}

func (e DictionaryErr) Error() string {
	return string(e)
}
