package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(s string) (string, error) {
	definitions, ok := d[s]
	if !ok {
		return "", errors.New("could not find the word you are looking for")
	}
	return definitions, nil
}
