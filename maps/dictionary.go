package dictionary

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	def, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return def, nil
}

//An interesting property of maps is that you can modify them without passing them as a pointer. This is because map is a reference type. Meaning it holds a reference to the underlying data structure, much like a pointer. The underlying data structure is a hash table, or hash map, and you can read more about hash tables here.
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
