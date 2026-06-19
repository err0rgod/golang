package main


import "fmt"

func main(){
	defer fmt.Print("\nDone exec")
	arr := [3]int{1, 2, 4}
	fmt.Print(arr[1])
}