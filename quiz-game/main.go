package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/eiannone/keyboard"
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

	var count int

	fmt.Println("Press enter to start the test:")
	_, key, err := keyboard.GetSingleKey()

	if err != nil {
		log.Println(err)
	}

	for key != 0x0D {
		_, key, err = keyboard.GetSingleKey()
		if err != nil {
			log.Println(err)
		}
	}

	timer := time.NewTimer(time.Second * 5)

	for i := 0; i < len(records); i++ {
		fmt.Println(records[i][0])
		answerCh := make(chan string)
		go func() {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			answerCh <- scanner.Text()
		}()
		select {
		case <-timer.C:
			fmt.Println("Number of correct answers is ", count)
			return
		case answer := <-answerCh:
			if answer == records[i][1] {
				count++
			}
		}
		fmt.Println("Number of correct answers is ", count)
	}
}
