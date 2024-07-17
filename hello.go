package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const helloSuffix = "!"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return helloPrefix(language) + name + helloSuffix
}

func helloPrefix(language string) string {
	prefix := englishHelloPrefix

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	}

	return prefix

}

func main() {
	fmt.Println(Hello("Omar", "English"))
}
