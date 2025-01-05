package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

// Read file content
func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}
	return data, nil
}

// Save binary data to a file
func saveToFile(filePath, data string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return fmt.Errorf("Error writing data to file: %v", err)
	}
	return nil
}

// Load the replacement map from a JSON file
func loadReplacementMap(mappingFilePath string) (map[string]string, error) {
	replacementMap := make(map[string]string)
	mappingData, err := ioutil.ReadFile(mappingFilePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading mapping file: %v", err)
	}

	err = json.Unmarshal(mappingData, &replacementMap)
	if err != nil {
		return nil, fmt.Errorf("Error parsing mapping file: %v", err)
	}

	return replacementMap, nil
}

// Apply replacement rules to the binary string
func replaceBinary(binaryStr string, replaceMap map[string]string) string {
	for old, new := range replaceMap {
		binaryStr = strings.ReplaceAll(binaryStr, old, new)
	}
	return binaryStr
}

func main() {
	// Specify file paths
	compressingFileExtension := "png"
	inputFilePath := "input_file." + compressingFileExtension
	outputFilePath := "compressed.txt"
	mappingFilePath := "chunk_mapping.json"

	// Step 1: Load the replacement map from chunk_mapping.json
	replacementMap, err := loadReplacementMap(mappingFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Map loaded")

	// Step 2: Read the input file (binary data)
	data, err := readFile(inputFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile("bytecode.txt", data, 0644)
	if err != nil {
		fmt.Printf("Error saving bytecode to file: %v\n", err)
		return
	}
	fmt.Println("Bite code saved")

	// Step 3: Convert the file content to binary string
	binaryStr := bytesToBinary(data)
	fmt.Println("Binary read")

	// Step 4: Apply replacement rules
	replacedBinary := replaceBinary(binaryStr, replacementMap)
	fmt.Println("Data compressed")

	// Save the replaced binary data to compressed.txt
	err = saveToFile(outputFilePath, replacedBinary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File compressed successfully!")
}
