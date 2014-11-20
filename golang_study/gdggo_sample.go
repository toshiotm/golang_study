package main

import "fmt"

func say(word string, ch chan int) {
	fmt.Println(word)
	ch <- 1
}


func main() {
	var ch chan int
	
	ch = make(chan int)

	go say("Hello World", ch)
	<-ch
}


