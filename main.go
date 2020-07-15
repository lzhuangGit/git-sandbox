package main

import "fmt"

func main() {
	fmt.Println(Greet())
	fmt.Printf(Greet_again())
}

func Greet() string {
	return "Hello, World!"
}

func Greet_again() string {
     return "Greetings, again!"
}