package main

import (
	"fmt"
)

const englishHelloPrefixx = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefixx + name
}

func main() {
	fmt.Println(Hello("World"))
}
