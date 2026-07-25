package main

import (
	"csv_to_json/csv_reader"
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

	fmt.Printf("Try to read %s\n", filePath)

	data, err := csv_reader.ReadCSV(filePath)

	if err != nil {
		fmt.Printf("Can't read file: %v\n", err)
		return
	}

	fmt.Println(data)
}
