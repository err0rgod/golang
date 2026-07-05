package main

import "fmt"

func main() {
	msgChan := make(chan string, 100)

	msgChan <- "hello, how are you"
	msgChan <- "i am fine"

	fmt.Println(<-msgChan)
	fmt.Println(<-msgChan)
}