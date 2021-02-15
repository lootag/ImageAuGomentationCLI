package commons

import (
	"encoding/csv"
	"os"
)

func ReadCsv(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		panic("Couldn't read the following annotation: " + filePath)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic("Couldn't extract records from the following annotation: " + filePath)
	}

	return records
}
