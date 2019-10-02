package dictionary

type Dictionary map[string]string
type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

var ErrNotFound = DictionaryError("could not find the word you were looking for")
var ErrWordExists = DictionaryError("word already present")
var ErrWordDoesNotExist = DictionaryError("word not present")

func (d Dictionary) Search(word string) (string, error) {
	def, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return def, nil
}

//An interesting property of maps is that you can modify them without passing them as a pointer.
//This is because map is a reference type. Meaning it holds a reference to the underlying data structure,
//much like a pointer. The underlying data structure is a hash table, or hash map,
func (d Dictionary) Add(word, definition string) error {
	_, found := d[word]
	if found {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, found := d[word]
	if !found {
		return ErrWordDoesNotExist
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
