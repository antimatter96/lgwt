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
