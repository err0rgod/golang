package main

import (
	"fmt"

)



func main() {
	taskmanager  := make(map[int]string)
	// create tasks in the map

	for true {
		option := 0
		fmt.Print("Enter you choice-----\n")
		fmt.Print("Enter 1 for adding a task.\n")
		fmt.Print("Enter 2 for completing a task.\n")
		fmt.Print("Enter 3 for listing all task.\n")
		fmt.Print("Enter 4 for exiting.\n")
		fmt.Scanf("%d", &option)
		switch option{
		case 1 :
			fmt.Println("Enter the task id and task. (with spaces)")
			var id int
			var task string
			
			fmt.Scanln(&id,&task)
			taskmanager[id] = task
		case 2:
			fmt.Print("Enter the task id to mark as completed.")
			var id int
			_,exists := taskmanager[id]
			if exists {
				delete(taskmanager,id)
			} else {
				fmt.Print("Invalid task ID.")
			}
		case 3 :
			fmt.Print("Listing all tasks ----------")
			for val,_ := range  taskmanager{
				fmt.Println(val)
			}
		case 4:
			return
		default : 
		fmt.Print("Select valid choice.")
		
		}
		

	}

}