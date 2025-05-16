package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func createFile(fileName string) string {
	_, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create the file: %v", err)
	}
	
	absolutePath, err := filepath.Abs(fileName)
	if err != nil {
		log.Fatalf("Failed to reach your file's path: %v", err)
	}

	return fmt.Sprintf("File '%v' created successfully at the directory %v\n", fileName, absolutePath)
}

func readFile(fileName string) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Failed to read your file: %v", err)
	}

	fmt.Println(string(fileContent))
}

func writeFile(fileName, data string) {
	err := os.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		log.Fatalf("Failed to write in your file: %v", err)
	}
}

func deleteFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		log.Fatalf("Failed to delete the file: %v", err)
	}
}