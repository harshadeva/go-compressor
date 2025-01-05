package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Convert bytes to binary string
func bytesToBinary(data []byte) string {
	var binaryStr strings.Builder
	for _, b := range data {
		binaryStr.WriteString(fmt.Sprintf("%08b", b))
	}
	return binaryStr.String()
}

// Function to convert file content to binary and append to txt file
func appendFileToTxt(filePath string, txtFilePath string) error {
	// Read file content
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error reading file %s: %v", filePath, err)
	}

	// Convert file content to binary
	binaryData := bytesToBinary(data)

	// Open txt file for appending
	txtFile, err := os.OpenFile(txtFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Error opening txt file for appending: %v", err)
	}
	defer txtFile.Close()

	// Append the binary data to the txt file
	_, err = txtFile.WriteString(binaryData + "\n")
	if err != nil {
		return fmt.Errorf("Error writing to txt file: %v", err)
	}

	return nil
}

// Function to process all files in the given directory and append their binary content to txt file
func processDirectory(dirPath string, txtFilePath string) error {
	// Read all files in the directory
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("Error reading directory: %v", err)
	}

	// Iterate through all files in the directory
	for _, file := range files {
		// Skip directories
		if file.IsDir() {
			continue
		}

		// Get the full path of the file
		filePath := filepath.Join(dirPath, file.Name())

		// Append the file's binary data to the txt file
		err := appendFileToTxt(filePath, txtFilePath)
		if err != nil {
			return fmt.Errorf("Error processing file %s: %v", file.Name(), err)
		}

		fmt.Printf("Processed file: %s\n", file.Name())
	}

	return nil
}

func main() {
	// Directory path (change this to your desired directory)
	dirPath := "files"
	// Output txt file path
	txtFilePath := "binary_file.txt"

	// Process the directory and append binary content to the txt file
	err := processDirectory(dirPath, txtFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Processing completed successfully!")
}
