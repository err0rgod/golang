package main

import (
	"fmt"

)

type task struct {
	id int
	task string
}


func main() {
	var taskmanager task[]
	// create tasks in the map

	for true {
		option := 0
		fmt.Print("Enter you choice-----\n")
		fmt.Print("Enter 1 for adding a task.\n")
		fmt.Print("Enter 2 for completing a task.\n")
		fmt.Print("Enter 3 for listing all task.\n")
		fmt.Print("Enter 4 for exiting.\n")
		fmt.Scan(&option)
		switch option{
		case 1 :
			fmt.Println("Enter the task id and task. (with spaces)")
			var id int
			var task string
			
			fmt.Scan(&id,&task)
			taskmanager.id = task
			
		case 2:
			fmt.Println("Enter the task id to mark as completed.")
			var id int
			fmt.Scan(&id)
			_,exists := taskmanager.id
			if exists {
				delete(taskmanager,id)
			} else {
				fmt.Println("Invalid task ID.")
			}
			
		case 3 :
			fmt.Println("Listing all tasks ----------")
			for val,tasks := range  taskmanager.task{
				fmt.Println(val,tasks)
			}
			
		case 4:
			return
			
		default : 
		fmt.Println("Select valid choice.")
		
		}
		

	}

}