package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func csvReader(filename string) [][]string {

	recordFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	reader := csv.NewReader(recordFile)

	records, _ := reader.ReadAll()

	// for _, row := range records {
	// 	fmt.Println(row[0])
	// 	fmt.Println(row[1])
	// }
	return records
}
