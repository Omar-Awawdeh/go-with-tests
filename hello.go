package main

import "fmt"

const (
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
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Omar", "English"))
}
