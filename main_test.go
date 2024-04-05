package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello World!"
	result := HelloWorld()
	if result != expected {
		t.Errorf("got %s, want %s", result, expected)
	}
}
