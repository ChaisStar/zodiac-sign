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
	var sheetSuffix string

	if response.Type == models.FrenchChinese {
		sheetSuffix = "_ch"
	} else {
		sheetSuffix = ""
	}

	builder.add(response, sheetSuffix)
}

func (builder ExcelFileBuilder) add(response models.Response, sheetSuffix string) {
	sheetName := response.Date.Format("2006-01-02") + sheetSuffix
	builder.createSheet(sheetName)

	column := string(rune('B' - 1 + response.Sign))
	var columnName string
	if response.Type == models.FrenchChinese {
		columnName = models.ChineseSign(response.Sign).String()
	} else {
		columnName = models.ZodiacSign(response.Sign).String()
	}

	builder.f.SetCellValue(sheetName, fmt.Sprintf("%s1", column), columnName)
	for label, text := range response.Texts {
		row := 2
		for row < 15 {
			value, err := builder.f.GetCellValue(sheetName, fmt.Sprintf("A%d", row))
			if err != nil {
				fmt.Println(err)
			}
			if value == "" {
				break
			}

			if value == label {
				break
			}
			row++
		}
		builder.f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), label)
		builder.f.SetCellValue(sheetName, fmt.Sprintf("%s%d", column, row), text)
	}
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

	sheet1Index := builder.f.GetSheetIndex("Sheet1")
	if sheet1Index >= 0 {
		builder.f.DeleteSheet("Sheet1")
	}
}

func (builder ExcelFileBuilder) Bytes() []byte {
	buffer, err := builder.f.WriteToBuffer()
	if err != nil {
		fmt.Println(err)
	}
	return buffer.Bytes()
}
