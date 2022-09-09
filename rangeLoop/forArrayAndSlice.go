package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c"}

	// with index and value
	fmt.Println("Both Index and Value")
	for i, letter := range letters {
		fmt.Printf("Index: %d Value: %s\n", i, letter)
		
	}

	// only value
	fmt.Println("\nOnly value")
	for _, letter := range letters {
		fmt.Printf("Value: %s\n", letter)
	}

	// Only index
	fmt.Println("\nOnly Index")
	for i := range letters {
		fmt.Printf("Index : %d\n", i)
	}

	// without index and value. Just print array value
	fmt.Println("\nWithout Index and Value")
	i := 0
	for range letters {
		fmt.Printf("Index : %d Value: %s\n", i, letters[i])
		i++
	}
}
