package main

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(dictionary, word, definition, t)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(dictionary, word, definition, t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete an existing key", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)
		assertError(t, err, nil)

		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)

	})

	t.Run("Trying to delete that which does not exist", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		err := dictionary.Delete(word)

		assertError(t, err, ErrDeleteNonExistingWord)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("testing update", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDef := "This is updated definition"
		dictionary.Update(word, newDef)

		assertDefinition(dictionary, word, newDef, t)
	})

	t.Run("testing adding new", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		newDef := "This is updated definition"
		err := dictionary.Update(word, newDef)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(dictionary Dictionary, testKey string, testVal string, t testing.TB) {

	t.Helper()
	got, err := dictionary.Search(testKey)

	if err != nil {
		t.Fatal("Should have found the Key", err)
	}

	assertStrings(t, got, testVal)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Unkown word", func(t *testing.T) {
		_, err := dictionary.Search("me")

		if err == nil {
			t.Fatal("Unknown word is supposed to throw error")
		}
		assertError(t, err, ErrNotFound)

	})
	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)

	})

}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("got %s wanted %s", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s wanted %s", got, want)
	}
}
