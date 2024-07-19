package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Omar", "English")
		want := "Hello, Omar!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in Spanish", func(t *testing.T) {
		got := Hello("Diego", "Spanish")
		want := "Hola, Diego!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in French", func(t *testing.T) {
		got := Hello("Louis", "French")
		want := "Bonjour, Louis!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q want: %q", got, want)
	}
}
