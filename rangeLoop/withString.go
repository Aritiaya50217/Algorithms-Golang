package main

import "fmt"

func main() {
	sample := "a£c"
	// "a" takes one byte as per UTF-8
	// "£" takes two byte as per UTF-8
	// "c" takes one byte as per UTF-8
	fmt.Printf("Length is %d\n", len(sample))

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%c\n", sample[i])
	}

	// with index and value
	fmt.Println("Both Index and Value")
	for i, letter := range sample {
		fmt.Printf("Start Index : %d Value: %s\n", i, string(letter))
	}

	// value
	fmt.Println("\nOnly value")
	for _, letter := range sample {
		fmt.Printf("Value : %s\n", string(letter))
	}

	// index
	fmt.Println("\nOnly Index")
	for i := range sample {
		fmt.Printf("Start Index: %d\n", i)
	}

}
