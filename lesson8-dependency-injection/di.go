package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s!", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func GreetTerminal() {
	Greet(os.Stdout, "Terminal")
}

func main() {
	GreetTerminal()

	handler := http.HandlerFunc(GreetHandler)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
