package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
)

const DataFileName = "data.json"

// Returns a path of the data-file
func GetDataFilePath() string {
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	return path.Join(workingDirectory, DataFileName)
}

// Creates a new data-file if it doesn't exist
func CreateDataFileIfNotExists() {
	filePath := GetDataFilePath()
	if _, err := os.Stat(filePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			file, err := os.Create(filePath)
			if err != nil {
				log.Fatalf("Error creating a data-file: %v", err)
			}
			defer file.Close()
			file.WriteString("{}")
		}
	}
}

// Saves given map as an actual map of tasks
func SaveTasksMap(tasks map[int]Task) {
	dataFilePath := GetDataFilePath()

	// Flags: open to write + create if not exists + remove content
	dataFile, err := os.OpenFile(dataFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Error opening data-file to write: %v", err)
	}
	defer dataFile.Close()

	encoder := json.NewEncoder(dataFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(&tasks); err != nil {
		log.Fatalf("Error encoding JSON to data-file: %v", err)
	}
}
