package Dictionary

import (
	"errors"
)

var (
	ErrorWordNotFound       = errors.New("Could not find the word you were looking for")
	ErrorWordAlreadyDefined = errors.New("Word already defined")
	ErrorWordNotDefined     = errors.New("Searched word is not defined")
)

type Dictionary map[string]string

func (dictionary Dictionary) Search(key string) (string, error) {
	searchResult, ok := dictionary[key]

	if !ok {
		return "", ErrorWordNotFound
	}

	return searchResult, nil
}

func (dictionary Dictionary) Add(key, value string) error {
	_, ok := dictionary[key]
	if ok {
		return ErrorWordAlreadyDefined
	}

	dictionary[key] = value
	return nil
}

func (dictionary Dictionary) Update(key, value string) error {
	_, ok := dictionary[key]
	if !ok {
		return ErrorWordNotDefined
	}

	dictionary[key] = value
	return nil
}

func (dictionary Dictionary) Delete(key string) {
	delete(dictionary, key)
}
