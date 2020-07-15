package main_test

import (
	"testing"

	main "."
)

func TestGreet(t *testing.T) {
	if main.Greet() != "Hello, World!" {
		t.Fail("failed first greeting test")
	}
	if main.Greet_again() != "Greetings, again!"
	   t.fail("failed greeting again test!!")
}
