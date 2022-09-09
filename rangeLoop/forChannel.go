package main

import "fmt"

func main() {
	ch := make(chan string)
	go pushToChannel(ch)
	for val := range ch {
		fmt.Println(val)
	}

	// []
	// push value to channel => [a , b , c]
	// for loop print a => [b,c]
	// for loop print b => [c]
	// for loop print c => []
	// close channel

}

func pushToChannel(ch chan<- string) {
	ch <- "a"
	ch <- "b"
	ch <- "c"
	close(ch)

}
