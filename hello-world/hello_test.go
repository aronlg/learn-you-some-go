package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Yoda")
	want := "Hello, Yoda"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
