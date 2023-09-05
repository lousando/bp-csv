package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No filepath provided.")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fileContents, error := os.ReadFile(filePath)

	if error != nil {
		log.Fatalf("%s is not a file.", filePath)
	}

	csvReader := csv.NewReader(strings.NewReader(string(fileContents)))

	// skip first line
	csvHeaders, error := csvReader.Read()

	// end of file
	if error == io.EOF {
		log.Fatalf("The file has no contents: %s", error)
	}

	if error != nil {
		log.Fatalf("ERROR: %s", error)
	}

	allRecords, error := csvReader.ReadAll()

	if error != nil {
		log.Fatalf("Failed to read contents of: %s", filePath)
	}

	for recordIndex := 0; recordIndex < len(allRecords); recordIndex++ {
		var record []string = allRecords[recordIndex]
		var rawTimestamp string = record[0]

		unixTimestamp, error := strconv.ParseInt(rawTimestamp, 10, 64)

		if error != nil {
			log.Fatalf("Failed to parse as unix timestamp: %s\n", rawTimestamp)
		}

		parsedTime := time.UnixMilli(unixTimestamp)

		record[0] = parsedTime.Format("01-02-2006 15:04")
	}

	outputCsvFilename := fmt.Sprintf("formatted_%s", strings.ToLower(path.Base(filePath)))
	outputCsvFile, error := os.Create(outputCsvFilename)

	if error != nil {
		log.Fatalf("Failed to create file: %s", outputCsvFilename)
	}

	var csvWriter = csv.NewWriter(outputCsvFile)

	csvWriter.Write(csvHeaders)
	csvWriter.WriteAll(allRecords)

	fmt.Printf("Formatted file written to: %s", outputCsvFilename)
}
