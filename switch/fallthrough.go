package main

import "fmt"

func notAllowed() {
	i := 45
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
		fallthrough
	case i < 50:
		fmt.Println("i is less than 50")
		fallthrough
		fmt.Println("Not allowed")
	case i < 100:
		fmt.Println("i is less than 100")
	}
	
}

func main() {
	i := 45
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
		// fallthrough
	case i < 50:
		fmt.Println("i is less than 50")
		fallthrough
	case i < 100:
		fmt.Println("i is less than 100")
	}

	// 45 < 50 => i is less than 50
	// fallthrough => 45 < 100 => i is less than 100

	notAllowed()
}
