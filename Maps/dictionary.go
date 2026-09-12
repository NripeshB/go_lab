package main

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("could not find the word you are looking for")

func (d Dictionary) Search(s string) (string, error) {
	definitions, ok := d[s]
	if !ok {
		return "", ErrorNotFound
	}
	return definitions, nil
}

func (d Dictionary) Add(key, val string) {
	d[key] = val
}
