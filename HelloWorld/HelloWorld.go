package main

import (
	"fmt"
)

const spanish = "Spanish"

const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "

func Hello(name, language string) string {
	var prefix string

	if name == "" {
		name = "World"
	}

	if language == spanish {
		prefix = spanishHelloPrefix
	} else {
		prefix = englishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
