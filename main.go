package main


import "fmt"

func main(){
	defer fmt.Print("Done exec")
	sum := 0
	for {
		sum +=1
		if (sum > 100) {break}
		fmt.Println(sum)
	}
}