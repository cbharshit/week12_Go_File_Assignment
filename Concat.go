package main

import "fmt"

func Concat(a string, b string) string {
	return a + b
}

func main() {
	str1 := "Hello"
	str2 := "World"
	result := Concat(str1, str2)
	fmt.Println("Concatenated String:", result)
}
