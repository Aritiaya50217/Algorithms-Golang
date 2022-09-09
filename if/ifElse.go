package main

import "fmt"

func ladder() {
	age := 29
	if age < 18 {
		fmt.Println("Kid")
	} else if age >= 18 && age < 40 {
		fmt.Println("Young")
	} else {
		fmt.Println("Old")
	}
}

func nested() {
	a := 1
	b := 2
	c := 3
	if a > b {
		if a > c {
			fmt.Println(" Biggest is a ")
		} else if b > c {
			fmt.Println(" Biggest is b ")
		}
	} else if b > c {
		fmt.Println("Biggest is b")
	} else {
		fmt.Println("Biggest is c")
	}

}
func main() {
	a := 6
	if a > 5 {
		fmt.Println("a is greater than 5")
	}

	b := 4
	if b > 3 && b < 6 {
		fmt.Println("a is within range")
	}

	c := 1
	d := 2
	if c > d {
		fmt.Println("c is greater than d")
	} else {
		fmt.Println("d is greater than c")
	}

	ladder()
	nested()
}
