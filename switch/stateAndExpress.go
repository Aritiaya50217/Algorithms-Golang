package main

import "fmt"

func main() {
	switch ch := "a"; ch {
	case "a":
		fmt.Println("a")
	case "b", "c":
		fmt.Println("b or c")
	default:
		fmt.Println("No matching character")
	}
}
