package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	a := 1 + 1

	fmt.Println(a)

	fmt.Println(Hello("world"))
}
