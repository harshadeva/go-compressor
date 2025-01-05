package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

// Convert binary string back to bytes
func binaryToBytes(binaryStr string) ([]byte, error) {
	var data []byte
	for i := 0; i < len(binaryStr); i += 8 {
		if i+8 > len(binaryStr) {
			break
		}
		byteValue, err := strconv.ParseUint(binaryStr[i:i+8], 2, 8)
		if err != nil {
			return nil, err
		}
		data = append(data, byte(byteValue))
	}
	return data, nil
}

// Check if the binary string is valid
func isValidBinary(binaryStr string) bool {
	for _, ch := range binaryStr {
		if ch != '0' && ch != '1' {
			return false
		}
	}
	return true
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

// Reverse the replacement rules to restore the original binary string
func reverseEncodeToBinary(replacedStr string, replaceMap map[string]string) string {
	for old, new := range replaceMap {
		replacedStr = strings.ReplaceAll(replacedStr, new, old)
	}
	return replacedStr
}

func main() {
	// Specify file paths
	compressingFileExtension := "jpg"
	inputFilePath := "input_file." + compressingFileExtension
	outputFilePath := "compressed.txt"
	restoredFilePath := "restored_file." + compressingFileExtension
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

	// Step 3: Convert the file content to binary string
	binaryStr := bytesToBinary(data)

	// Save the binary data to binary_file.txt
	err = saveToFile("binary_file.txt", binaryStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Binary file saved")

	// Step 4: Apply replacement rules
	replacedBinary := replaceBinary(binaryStr, replacementMap)
	fmt.Println("Data compressed")

	// Save the replaced binary data to compressed.txt
	err = saveToFile(outputFilePath, replacedBinary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Compressed file saved")

	// Step 5: Reverse the process to restore the original binary
	restoredBinary := reverseEncodeToBinary(replacedBinary, replacementMap)

	// valid := isValidBinary(restoredBinary)
	// if !valid {
	// 	fmt.Printf("Binary validation failed")
	// 	return
	// }

	// Step 6: Convert the restored binary back to bytes
	restoredData, err := binaryToBytes(restoredBinary)
	if err != nil {
		fmt.Printf("Error converting binary to bytes: %v\n", err)
		return
	}

	// Step 7: Save the restored data to a new file
	err = saveToFile(restoredFilePath, string(restoredData))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Process completed successfully!")
}
