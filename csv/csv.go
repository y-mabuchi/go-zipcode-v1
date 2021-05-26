package csv

import (
	"encoding/csv"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type Csv struct {
	filename string
}

func NewCsv(filename string) *Csv {
	return &Csv{filename: filename}
}

func (c *Csv) GetFilename() string {
	return c.filename
}

func (c *Csv) Read() [][]string {
	// Read file as ShiftJIS
	sjisFile, err := os.Open(c.filename)
	// Typical error handling in Go
	if err != nil {
		log.Fatalf("action=Csv.Read.os.Open, status=error: %v\n", err)
	}
	// defer run after the command when this function done
	defer sjisFile.Close()

	// Convert to UTF-8
	utf8reader := transform.NewReader(sjisFile, japanese.ShiftJIS.NewDecoder())

	reader := csv.NewReader(utf8reader)
	var multiLine [][]string

	for {
		line, err := reader.Read()
		if err != nil {
			log.Fatalf("action=Csv.Read.reader.Read, status=error: %v\n", err)
		}
		multiLine = append(multiLine, line)
	}

	return multiLine
}
