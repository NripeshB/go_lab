package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Spanish", "Chris")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)

	})
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("English", "Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("German", "Anna")
		want := "Halo, Anna"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
