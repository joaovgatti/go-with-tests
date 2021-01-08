package main

import "testing"


func TestHelloWorld(t *testing.T){

	got := HelloWorld()
	want := "Hello World"

	if got != want {
		t.Errorf("got %q want %q",got ,want)
	}


}

func TestHelloUser(t *testing.T) {
	got := HelloUser("joao")
	want := "Hello Joao"

	if got != want {
		t.Errorf("got %q want %q",got,want)
	}
}
