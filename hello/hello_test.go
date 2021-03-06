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
		got := HelloUser("Joao","")
		want := "Hello, Joao"
		assertCorrectMessage(t,got,want)

	})

	t.Run("say hello when an empty string is supplied", func (t *testing.T){
		got := HelloUser("","")
		want := "Hello, World"
		assertCorrectMessage(t,got,want)
	})

	t.Run("in Spanish", func (t *testing.T){
		got := HelloUser("De la Mancha","spanish")
		want := "Hola, De la Mancha"
		assertCorrectMessage(t,got,want)
	})

	t.Run("in Portuguese", func (t *testing.T){
		got := HelloUser("Machado","portuguese")
		want := "Olá, Machado"
		assertCorrectMessage(t,got,want)
	})

}
