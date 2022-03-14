package main

import "fmt"

func Hello(name string) string {
	const helloPrefix = "Hello, "
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("Yoda"))
}
