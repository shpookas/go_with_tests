package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", " ")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", " ")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Elae", "French")
		want := "Bonjour, Elae"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Lithuanian", func(t *testing.T) {
		got := Hello("Vytautas", "Lithuanian")
		want := "Labas, Vytautas"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	//t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
