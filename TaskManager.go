package TaskManager

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Id   int
	Name string
	Done bool
}

func LoadTasks() []Task {
	byte, err := os.ReadFile("task.json")
	if err != nil {
		fmt.Println("Error Occured while reding file. ", err)
		return nil
	}

	var taskslice []Task
	if err := json.Unmarshal(byte, &taskslice); err != nil {
		fmt.Println("Error while parsing json.")
		return nil
	}
	return taskslice
}

func SaveTasks(taskslice []Task) {
	byte, err := json.Marshal(taskslice)
	if err != nil {
		fmt.Println("Error while parsing Struct to Json ")
		return
	}
	os.WriteFile("task.json", byte, 0644)
}

func main() {
	// create tasks in the map
	tasks := LoadTasks()

	for {
		option := 0
		fmt.Print("\n 1. Add Task 2. Complete Task 3. List Tasks 4. Quit\n ")
		fmt.Scan(&option)
		switch option {
		case 1:
			fmt.Println("Enter the task id and task. (with spaces)")
			var id int
			var task string

			fmt.Scan(&id, &task)
			tasks = append(tasks, Task{Id: id, Name: task})
			SaveTasks(tasks)

		case 2:
			fmt.Println("Enter the task id to mark as completed.")
			var id int
			fmt.Scan(&id)
			found := false
			for i, t := range tasks {
				if t.Id == id {
					tasks[i].Done = true
					found = true
					fmt.Println("Marked done. ")
					SaveTasks(tasks)
					break
				}
			}
			if !found {
				fmt.Println("Invalid ID.")
			}

		case 3:
			fmt.Println("Listing all tasks ----------")
			for _, t := range tasks {
				fmt.Printf("[%d] -> %s -isDone %v\n", t.Id, t.Name, t.Done)
			}

		case 4:
			return

		default:
			fmt.Println("Invalid choice.")

		}

	}

}
