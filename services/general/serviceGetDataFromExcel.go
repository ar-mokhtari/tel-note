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
		result = append(result, record)
	}
	if hasHeader {
		result = result[1:]
	}
	return result, nil
}
