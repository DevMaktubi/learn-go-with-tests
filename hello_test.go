package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("sayig hellow to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("","")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Breno", "Spanish")
		want := "Hola, Breno!"
		assertCorrectMessage(t,got,want)
	}) 
	t.Run("In French", func(t *testing.T) {
		got := Hello("Breno", "French")
		want := "Bonjour, Breno!"
		assertCorrectMessage(t,got,want)
	}) 
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}