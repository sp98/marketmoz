package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CSVReader(filePath string) ([][]string, error) {

	// Open the file
	recordFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("failed to open file", err)
		return nil, err
	}

	// Setup the reader
	reader := csv.NewReader(recordFile)

	// Read the records
	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println("failed to read from csv file", err)
		return nil, err
	}

	err = recordFile.Close()
	if err != nil {
		fmt.Println("failed to close file", err)
		return nil, err
	}

	return allRecords, nil

}
