package main

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("testKey", "A test key ")
	want := "A test key "
	got, err := dictionary.Search("testKey")
	if err != nil {
		t.Fatal("Should have found the Key", err)
	}

	assertStrings(t, got, want)

}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Unkown word", func(t *testing.T) {
		_, err := dictionary.Search("me")

		if err == nil {
			t.Fatal("Unknown word is supposed to throw error")
		}
		assertError(t, err, ErrorNotFound)

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
