package main

import (
	"csv_to_json/converter"
	"csv_to_json/reader"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Error: The path to CSV file not found!")
		fmt.Println("Uses: go run main.go path_to_file")
		return
	}

	filePath := os.Args[1]

	fmt.Printf("Attempt to read the file %s\n", filePath)

	data, err := reader.ReadCSV(filePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Attempt to convert to JSON...")
	jsonString, err := converter.CSVToJSON(data)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't convert: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(jsonString)
}
