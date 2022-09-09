package main

import "fmt"

func main() {
	printType("test_string")
	printType(2)
	printType(true)
}

func printType(t interface{}) {
	switch v := t.(type) {
	case string:
		fmt.Println("Type : string")
	case int:
		fmt.Println("Type : int")
	case bool:
		fmt.Println("Type : boolean")
	default:
		fmt.Printf("Unknown Type : %T", v)
	}
}
