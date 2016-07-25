package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("violations.csv")
	// check file
        if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// close f
	defer f.Close()

	csvr := csv.NewReader(bufio.NewReader(f))
        lnc := 0

	for {
		record, err := csvr.Read()
		if err == io.EOF {
			break
		}
	   // fmt.Println(record)
	   fmt.Println("Row:", lnc, " had:  ", record, " with ",  len(record), " items")

		for value := range record {
			fmt.Printf(" %v\n", record[value])
		}
	   fmt.Println()
           lnc += 1
	}
}
