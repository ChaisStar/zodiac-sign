package services

import (
	"archive/zip"
	"bytes"
	"fmt"
	"log"
)

func CreateZip(files []bytes.Buffer) []byte {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)
	for i, file := range files {
		zipFile, err := zipWriter.Create(fmt.Sprintf("%d.xlsx", i))
		if err != nil {
			fmt.Println(err)
		}

		log.Printf("%s", file.Bytes())

		_, err = zipFile.Write(file.Bytes())
		if err != nil {
			fmt.Println(err)
		}
	}

	err := zipWriter.Close()
	if err != nil {
		fmt.Println(err)
	}

	return buf.Bytes()
}
