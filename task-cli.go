package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Adds a new task to the map of tasks
func AddTask(description string) {
	tasks := GetTasksMap()

	var newId int
	for id := range tasks {
		newId = max(newId, id)
	}
	newId++

	newTask := NewTask(description)
	tasks[newId] = *newTask
	SaveTasksMap(tasks)
}

// Updates description of a task by its id
func UpdateTaskDescription(id int, description string) {
	tasks := GetTasksMap()

	if targetTask, ok := tasks[id]; ok {
		targetTask.Description = description
		targetTask.UpdatedAt = time.Now()
		tasks[id] = targetTask
		SaveTasksMap(tasks)
	} else {
		log.Fatal("Error updating non-existent task")
	}
}

// Deletes task from tasks map by its id
func DeleteTask(id int) {
	tasks := GetTasksMap()

	if _, ok := tasks[id]; ok {
		delete(tasks, id)
		SaveTasksMap(tasks)
	} else {
		log.Fatal("Error deleting non-existent task")
	}
}

// Updates status of a task by its id
func UpdateTaskStatus(id int, newStatus string) {
	tasks := GetTasksMap()

	if targetTask, ok := tasks[id]; ok {
		targetTask.Status = newStatus
		targetTask.UpdatedAt = time.Now()
		tasks[id] = targetTask
		SaveTasksMap(tasks)
	} else {
		log.Fatal("Error updating status of non-existent task")
	}
}

// Clears a map of tasks
func ClearTasksMap() {
	emptyMap := make(map[int]Task)
	SaveTasksMap(emptyMap)
}

func init() {
	CreateDataFileIfNotExists()
}

func main() {
	argsCount := len(os.Args) - 1
	actionType := os.Args[1]

	switch actionType {
	case "help":
		fmt.Println(syntax)

	case "add":
		if argsCount != 2 {
			log.Fatal("Error. Wrong using of add. Try \"task-cli help\"")
		}
		description := os.Args[2]
		AddTask(description)

	case "update":
		if argsCount != 3 {
			log.Fatal("Error. Wrong using of update. Try \"task-cli help\"")
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Error. Wrong task id format: %v\n. Try \"task-cli help\"", err)
		}
		description := os.Args[3]
		UpdateTaskDescription(id, description)

	case "delete":
		if argsCount != 2 {
			log.Fatal("Error. Wrong using of delete. Try \"task-cli help\"")
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Error. Wrong task id format: %v\n. Try \"task-cli help\"", err)
		}
		DeleteTask(id)

	case "mark-todo":
		if argsCount != 2 {
			log.Fatal("Error. Wrong using of mark-todo. Try \"task-cli help\"")
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Error. Wrong task id format: %v\n. Try \"task-cli help\"", err)
		}
		UpdateTaskStatus(id, "todo")

	case "mark-in-progress":
		if argsCount != 2 {
			log.Fatal("Error. Wrong using of mark-in-progress. Try \"task-cli help\"")
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Error. Wrong task id format: %v\n. Try \"task-cli help\"", err)
		}
		UpdateTaskStatus(id, "in-progress")

	case "mark-done":
		if argsCount != 2 {
			log.Fatal("Error. Wrong using of mark-done. Try \"task-cli help\"")
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Error. Wrong task id format: %v\n. Try \"task-cli help\"", err)
		}
		UpdateTaskStatus(id, "done")

	case "list":
		if argsCount == 1 {
			tasksMap := GetTasksMap()
			fmt.Println(GetJSONStringFromMap(tasksMap))
		} else if argsCount == 2 {
			status := os.Args[2]
			if status != "todo" && status != "in-progress" && status != "done" {
				log.Fatal("Error. Wrong unknown status. Try \"task-cli help\"")
			}
			tasksMap := GetTasksMapByStatus(os.Args[2])
			fmt.Println(GetJSONStringFromMap(tasksMap))
		} else {
			log.Fatal("Error. Wrong using of list. Try \"task-cli help\"")
		}

	case "clear":
		if argsCount == 1 {
			ClearTasksMap()
		} else {
			log.Fatal("Error. Wrong using of clear. Try \"task-cli help\"")
		}

	default:
		log.Fatal("Error. Unknown action. Try \"task-cli help\"")
	}
}
