package general

import (
	"encoding/csv"
	"io"
	"os"
)

func GetDataFromExcel(path string, hasHeader bool) (result [][]string, err error) {
	file, _ := os.Open(path)
	// parse as csv
	reader := csv.NewReader(file)
	var results [][]string
	for {
		// read one row from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// add record to result set
		results = append(results, record)
	}
	if hasHeader {
		results = results[1:]
	}
	return results, nil
}
