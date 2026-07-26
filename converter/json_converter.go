package converter

import (
	"encoding/json"
	"fmt"
)

func CSVToJSON(records [][]string) (string, error) {
	if len(records) == 0 {
		return "", fmt.Errorf("CSV data are empty")
	}

	headers := records[0]

	var result []map[string]string

	for i := 1; i < len(records); i++ {
		row := records[i]

		if len(row) != len(headers) {
			return "", fmt.Errorf("Row %d has %d columns, but expecting %d.", i, len(row), len(headers))
		}

		rowMap := make(map[string]string)

		for j := 0; j < len(row); j++ {
			key := headers[j]
			value := row[j]
			rowMap[key] = value
		}

		result = append(result, rowMap)
	}

	jsonBytes, err := json.Marshal(result)

	if err != nil {
		return "", fmt.Errorf("JSON marshaling error:: %v", err)
	}

	return string(jsonBytes), nil
}
