package main

import (
	"fmt"
)

func main() {
	taskmanager  := make(map[int]string)
	taskmanager[1] ="do dishes"
	taskmanager[2] = "wash clothes"

	for id,t := range taskmanager{
		fmt.Printf("Task %d : %s\n",id,t)
	}
}