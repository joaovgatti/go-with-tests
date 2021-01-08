package main

import (
	"testing"
)


func TestHelloUser(t *testing.T) {


	assertCorrectMessage := func (t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q",got,want)
		}
	}

	t.Run("saying hello to people", func (t *testing.T){
		got := HelloUser("Joao")
		want := "Hello, Joao"
		assertCorrectMessage(t,got,want)

	})

	t.Run("say hello when an empty string is supplied", func (t *testing.T){
		got := HelloUser("")
		want := "Hello, World"
		assertCorrectMessage(t,got,want)
	})
}
