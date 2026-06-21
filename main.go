package main

import "fmt"


type Shape interface{
	Area() float64
}

type Rectangle struct{
	height float64
	width float64

}

type Circle struct{
	radius float64
}

type Triangle struct{
	sidelength float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}
func (c Circle) Area() float64 {
	return c.radius* 2//for test using 2 as stale formula  
}
func (t Triangle) Area() float64{
	return t.sidelength * 3
}


func printArea(s Shape) {
	fmt.Print(s.Area(),"\n")
}

func main() {
	r := Rectangle{10,10}
	c:= Circle{10}
	t:= Triangle{20}
	printArea(r)
	printArea(t)
	printArea(c)
}