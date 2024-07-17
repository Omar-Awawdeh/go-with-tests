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

func helloPrefix(language string) string {
	switch language {
	case "Spanish":
		return spanishHelloPrefix
	case "French":
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("Omar", "English"))
}
