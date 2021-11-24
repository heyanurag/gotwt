package util

import (
	"fmt"
	"gotwt/twitter"

	"github.com/xuri/excelize/v2"
)

func saveToExcel(pool []twitter.User) {

	file := excelize.NewFile()

	index := file.NewSheet("Sheet2")

	// Set value of a cell.
	for i := 0; i < len(pool); i++ {
		file.SetCellValue("Sheet2", fmt.Sprintf("A%d", i+1), pool[i].Id)
		file.SetCellValue("Sheet2", fmt.Sprintf("B%d", i+1), pool[i].Name)
		file.SetCellValue("Sheet2", fmt.Sprintf("C%d", i+1), pool[i].Username)
	}

	// Set active sheet of the workbook.
	file.SetActiveSheet(index)

	// Save spreadsheet by the given path.
	if err := file.SaveAs("excel/logFile.xlsx"); err != nil {
		fmt.Println(err)
	}
}
