package main

import "fmt"

const englishGreeting = "Hello, "
const spanishGreeting = "Hola, "
const germanGreeting = "Halo, "

func Hello(lang, name string) string {
	if name == "" {
		name = "World"
	}
	return getPref(lang) + name

}

func getPref(lang string) (prefix string) {

	switch lang {
	case "Spanish":
		prefix = spanishGreeting
	case "German":
		prefix = germanGreeting
	default:
		prefix = englishGreeting
	}

	return
}

func main() {
	fmt.Println(Hello("Enlgish", "World"))

}
