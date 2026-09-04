package main

import "fmt"

const englishGreeting = "Hello, "

func Hello(name string) string {
	if name == "" {
		return englishGreeting + "World"
	}

	return englishGreeting + name
}

func main() {
	fmt.Println(Hello("World"))

}
