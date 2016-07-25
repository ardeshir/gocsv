package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("violations.csv")

	csvr := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := csvr.Read()
		if err == io.EOF {
			break
		}
	   fmt.Println(record)
	   fmt.Println(len(record))
		for value := range record {
			fmt.Printf(" %v\n", record[value])
		}
	}
}
