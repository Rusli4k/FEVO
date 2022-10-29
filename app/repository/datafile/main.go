package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const filePath = "example.csv"

// Func ReadCsvFile takes named file, open it,
// read it and return slice of slices with lines
// from file and error if something gone wrong.
func ReadCsvFile(s string) ([][]string, error) {
	f, err := os.Open(s)

	if err != nil {
		return nil, fmt.Errorf("Unable to read input file "+s, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Unable to read input file "+s, err)
	}

	return records, nil
}

func main() {
	records, _ := ReadCsvFile(filePath)
	for _, v := range records {
		fmt.Println(v)
	}
	// fmt.Println(records)
}
