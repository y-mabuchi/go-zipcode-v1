package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

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
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("action=Csv.Read.reader.Read, status=error: %v\n", err)
		}
		multiLine = append(multiLine, line)
	}

	return multiLine
}

func (c *Csv) Write(data [][]string) {
	// Create timestamp for file name
	format := "20060102150405"
	timestamp := time.Now().Format(format)
	fmt.Println(timestamp)

	// Create file name
	filename := timestamp + "_filtered_zipcode.csv"

	// Get current directory
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("action=Csv.Write.Getwd, status=error: %v\n", err)
	}

	// Create path
	path := filepath.Join(pwd, "result_files", filename)

	// Create empty file
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("action=os.Create, status=error: %v\n", err)
	}

	// Call writer
	writer := csv.NewWriter(file)
	defer func() {
		writer.Flush()
		file.Close()
	}()

	// Write data
	for _, v := range data {
		err := writer.Write(v)
		if err != nil {
			log.Fatalf("action=writer.Write, status=error: %v\n", err)
		}
	}
}
