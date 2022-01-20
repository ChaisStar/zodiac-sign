package services

import (
	"asia/models"
	"bytes"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func CreateFile(response models.Response) *bytes.Buffer {
	f := excelize.NewFile()
	index := f.NewSheet("Sheet")
	f.SetCellValue("Sheet", "A1", "Date")
	f.SetCellValue("Sheet", "B1", "Sign")
	f.SetCellValue("Sheet", "C1", "Text")
	f.SetCellValue("Sheet", "A1", response.Date.Format("2006-01-02"))
	f.SetCellValue("Sheet", "B1", response.Sign)
	f.SetCellValue("Sheet", "C1", response.Text)

	f.SetActiveSheet(index)

	buffer, err := f.WriteToBuffer()
	if err != nil {
		fmt.Println(err)
	}

	return buffer
	// return bytes.NewBufferString(fmt.Sprintf("Date: %s, Sign: %d, Text: %s", response.Date.Format("2006-01-02"), response.Sign, response.Text))
}
