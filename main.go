package main

import (
	"fmt"
)

type digit float64

type Rectangle struct {
	height digit
	width digit
}


func (r Rectangle) Area() digit {
	return r.height * r.width
}

func main(){
	u := Rectangle {
		height: 10,
		width: 20,
	}
	fmt.Print(u.Area())
}