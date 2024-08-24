package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func readCsvFile(filePath string) [][]string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	var records [][]string
	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		records = append(records, record)
	}
	return records
}
func main() {
	records := readCsvFile("default.csv")
	fmt.Println(records)
}
