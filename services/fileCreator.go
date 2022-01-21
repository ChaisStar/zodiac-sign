package services

import (
	"fmt"

	"github.com/ChaisStar/zodiac-sign/models"

	"github.com/xuri/excelize/v2"
)

type ExcelFileBuilder struct {
	f *excelize.File
}

func NewBuilder() ExcelFileBuilder {
	builder := new(ExcelFileBuilder)
	builder.f = excelize.NewFile()
	return *builder
}

func (builder ExcelFileBuilder) Add(response models.Response) {
	sheetName := response.Date.Format("2006-01-02")
	builder.createSheet(sheetName)

	column := string(rune('A' - 1 + response.Sign))
	builder.f.SetCellValue(sheetName, fmt.Sprintf("%s1", column), response.Sign.String())
	builder.f.SetCellValue(sheetName, fmt.Sprintf("%s2", column), response.Text)
}

func (builder ExcelFileBuilder) createSheet(sheetName string) {
	sheets := builder.f.GetSheetList()
	containsSheet := false
	for _, sheet := range sheets {
		if sheet == sheetName {
			containsSheet = true
		}
	}

	if !containsSheet {
		builder.f.NewSheet(sheetName)
	}
}

func (builder ExcelFileBuilder) Bytes() []byte {
	buffer, err := builder.f.WriteToBuffer()
	if err != nil {
		fmt.Println(err)
	}
	return buffer.Bytes()
}
