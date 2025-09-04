package main

import (
	"fmt"
)

// https://go.dev/tour/methods/1

func main() {
	m1 := make(map[string]string)

	m1["name"] = "Alice"
	m1["course"] = "golang"
	m1["site"] = "example.com"

	fmt.Println(m1)

	if v, ok := m1["name"]; ok {
		fmt.Println("name:", v)
	} else {
		fmt.Println("no name")
	}

	if v, ok := m1["major"]; ok {
		fmt.Println("major:", v)
	} else {
		fmt.Println("no major")
	}
}
