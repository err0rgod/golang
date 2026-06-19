package main

import (
	"fmt"
)

type Rectangle struct {
	height float64
	width float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func main(){
	u := Rectangle {
		height: 10,
		width: 20,
	}
	fmt.Print(u.Area())
}