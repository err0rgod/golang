package main

import "fmt"

func main() {
	messageChan := make(chan string)

	messageChan <- "ping"

	msg := <-messageChan

	fmt.Print(msg)
}