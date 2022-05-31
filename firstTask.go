package main

import "fmt"

func main() {
	str := "Hello"
	var count int
	for range str {
		count++
	}
	fmt.Println(count)
}
