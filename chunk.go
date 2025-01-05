package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Function to generate the next short code, avoiding '1' and '0'
func generateShortCode(index int) string {
	availableDigits := []rune{'2', '3', '4', '5', '6', '7', '8', '9'}
	code := ""

	for index >= 0 {
		remainder := index % len(availableDigits)
		code = string(availableDigits[remainder]) + code
		index = index/len(availableDigits) - 1
	}

	return "@" + code + "*"
}

// Function to divide the binary string into 144-bit chunks
func chunkBinaryString(binaryString string, chunkSize int) map[string]string {
	chunkMapping := make(map[string]string)
	shortCodeIndex := 0

	// Iterate through the binary string and create chunks
	for i := 0; i < len(binaryString); i += chunkSize {
		end := i + chunkSize
		if end > len(binaryString) {
			end = len(binaryString)
		}

		chunk := binaryString[i:end]
		shortCode := generateShortCode(shortCodeIndex)
		chunkMapping[chunk] = shortCode
		shortCodeIndex++
	}

	return chunkMapping
}

// Save the mapping to a JSON file
func saveMappingToFile(chunkMapping map[string]string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	return encoder.Encode(chunkMapping)
}

func main() {
	// Input file path
	inputFilePath := "binary_file.txt"
	outputFilePath := "chunk_mapping.json"

	// Read binary string from file
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	binaryString := string(data)

	// Define chunk size (144 bits or 18 bytes)
	chunkSize := 144

	// Generate chunks and assign short codes
	chunkMapping := chunkBinaryString(binaryString, chunkSize)

	// Save the mapping to a JSON file
	err = saveMappingToFile(chunkMapping, outputFilePath)
	if err != nil {
		fmt.Println("Error saving chunk mapping:", err)
		return
	}

	fmt.Println("Chunk mapping saved to", outputFilePath)
}
