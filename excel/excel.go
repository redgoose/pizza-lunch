package excel

import (
	"fmt"
	"regexp"

	"github.com/xuri/excelize/v2"
)

func ProcessFile(fileName string) ([][]string, error) {
	processedRows := [][]string{}

	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s", fileName)
	}

	firstSheet := f.WorkBook.Sheets.Sheet[0].Name

	rows, err := f.GetRows(firstSheet)
	if err != nil {
		return nil, fmt.Errorf("failed to get rows")
	}

	re := regexp.MustCompile(`^\d+$`)

	for _, row := range rows {
		// skip rows that dont have a student number
		if !(len(row) >= 2 && re.MatchString(row[1])) {
			continue
		}

		processedRow := []string{}
		for _, colCell := range row {
			// skip empty cells
			if colCell == "" {
				continue
			}
			processedRow = append(processedRow, colCell)
		}
		processedRows = append(processedRows, processedRow)
	}

	return processedRows, nil
}
