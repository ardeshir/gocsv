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

	csvr = csv.NewReader(budio.NewReader(f))
	for {
		record, err := csvr.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
		fmt.Println(len(recoord))
		for vlaue := range record {
			fmt.Printf(" %v\n", record[value])
		}
	}
}
