package main

type Dictionary map[string]string

const (
	ErrorNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrorWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {

	definition, ok := d[key]

	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil

}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) {

	d[word] = definition

}
