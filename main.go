package main

import (
	"fmt"
)

func add(a , b int) (x,y int) {
	x = a+b
	y= (a-b)
	return
}

func main() {
	a,b := add(10,20)
	fmt.Println(a,b)
}