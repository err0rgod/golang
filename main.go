package main

import (
	"fmt"
)

func main() {
	var name string
	var age int

	fmt.Print("Enter name and age : ")
	_,err := fmt.Scanf("%s %d",&name,&age)
	if err != nil {
		fmt.Print("Error reading input.")
		return
	}
	fmt.Printf("Name : %s , age : %d ", name, age)
}