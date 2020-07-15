package main_test

import (
	"testing"

	main "."
)

func TestGreet(t *testing.T) {
	if main.Greet() != "Hello, World!" {
		t.Fail()
	}
	if main.Greet_again() != "Greetings, again!" {
	   t.Fail()
	}
}
