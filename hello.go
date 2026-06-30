package main

import "fmt"

func main() {
	fmt.Println(Hello("Juan"))
}

func Hello(name string) string {
	return "Hello, " + name
}
