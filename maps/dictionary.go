package main

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(key string) (string, error) {

	definition, ok := d[key]

	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil

}

func (d Dictionary) Add(key, definition string) {
	d[key] = definition
}