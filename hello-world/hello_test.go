package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying 'hello' to people", func(t *testing.T) {
		got := Hello("Anagh", "")
		want := "Hello Anagh"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty name defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Anagh", "Spanish")
		want := "Hola Anagh"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Anagh", "French")
		want := "Bonjour Anagh"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
