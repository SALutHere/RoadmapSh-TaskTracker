package main

import (
	"encoding/json"
	"log"
	"os"
)

const syntax = `Usage: task-cli [ACTION] [OPTIONS]...
Allows you to manage your tasks.

Actions with their options:
	help				gives you a syntax-hint

	add DESCRIPTION			adds a new task with specified description
	update ID DESCRIPTION		updates a description of task by the specified id
	delete ID			deletes a task by the specified id
	
	mark-todo ID			sets status "todo" to the task by specified id
	mark-in-progress ID		sets status "in-progress" to the task by specified id
	mark-done ID			sets status "done" to the task by specified id
	
	list [STATUS]			lists your tasks. If you specified a status, lists
					your tasks only with specified status
						allowed statuses: "todo", "in-progress", "done"
									
	clear				clears your tasks list`

// Returns an actual map of tasks
func GetTasksMap() map[int]Task {
	dataFilePath := GetDataFilePath()
	dataFile, err := os.Open(dataFilePath)
	if err != nil {
		log.Fatalf("Error opening data-file to read: %v", err)
	}
	defer dataFile.Close()

	decoder := json.NewDecoder(dataFile)
	tasks := make(map[int]Task)
	if err := decoder.Decode(&tasks); err != nil {
		log.Fatalf("Error decoding JSON from data-file: %v", err)
	}

	return tasks
}

// Returns tasks with specified status from an actual map of tasks
func GetTasksMapByStatus(status string) map[int]Task {
	tasks := GetTasksMap()
	resultTasksMap := make(map[int]Task)
	for id, task := range tasks {
		if task.Status == status {
			resultTasksMap[id] = task
		}
	}
	return resultTasksMap
}

// Returns a string with prettified JSON of tasks
func GetJSONStringFromMap(tasks map[int]Task) string {
	jsonString, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling map to JSON: %v", err)
	}
	return string(jsonString)
}
