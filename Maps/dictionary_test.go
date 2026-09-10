package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Unkown word", func(t *testing.T) {
		_, err := dictionary.Search("me")
		want := "could not find the word you are looking for"

		if err == nil {
			t.Fatal("Unknown word is supposed to throw error")
		}
		assertStrings(t, err.Error(), want)

	})
	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)

	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s wanted %s", got, want)
	}
}
