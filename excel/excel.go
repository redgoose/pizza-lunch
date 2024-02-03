package excel

import (
	"fmt"
	"regexp"

	"github.com/xuri/excelize/v2"
)

func ProcessFile(fileName string, sheetName string) ([][]string, error) {
	processedRows := [][]string{}

	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to open file: %s", fileName)
	}
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to find sheet: %s", sheetName)
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
