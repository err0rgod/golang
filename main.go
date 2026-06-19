package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func isAdult(u User) bool {
	if(u.Age >= 18) {return true}
	return false
}

func main(){
	u := User{"alice",9}
	v := User{"mdhmasale",19}
	if (isAdult(u)) {fmt.Print("Chota bacha\n")
	} else {fmt.Print("Bada Bachn\n")};
	if (isAdult(v)) {fmt.Print("Chota Bacha\n")
	} else {fmt.Print("Bada Bachn\n")}
}