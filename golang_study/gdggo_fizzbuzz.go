package main

import "fmt"
import "strconv"

func fizzbuzz(in chan int, out chan string) {
	
	for {
		var n int = <- in

		switch {
		case n%15 == 0: out <- "fizzbuzz"
		case n%3 == 0: out <-"fizz" 
		case n%5 == 0: out <- "buzz"
		default:	out <- strconv.Itoa(n)
		}
	}
}

func main() {
	ch := make(chan int)
	out := make(chan string)
	go fizzbuzz(ch, out)
	for i:=1; i<100; i++ {
		ch <-i
		s := <-out
		fmt.Println(s)
	} 
}
