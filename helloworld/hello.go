package main

import (
	"fmt"
	"strings"
)

const (
	spanish = "spanish"
	french  = "french"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	helloSuffix        = "!"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return helloPrefix(language) + name + helloSuffix
}

func helloPrefix(language string) (prefix string) {
	switch strings.ToLower(language) {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Omar", "English"))
}
