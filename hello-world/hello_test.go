package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Scott", "")
		want := "Hello, Scott"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In French", func(t *testing.T) {
		got := Hello("Jean", "French")
		want := "Bonjour, Jean"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In Korean", func(t *testing.T) {
		got := Hello("Lee", "Korean")
		want := "Annyeong, Lee"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
