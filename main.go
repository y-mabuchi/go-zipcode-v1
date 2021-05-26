package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/y-mabuchi/go-zipcode-v1/csv"
)

func init() {
	log.SetPrefix("Zipcode Converter: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	filename := flag.String("filename", "zipcode.csv", "CSV file name of Zipcode.")
	flag.Parse()
	fmt.Printf("file name: %v\n", *filename)

	c := csv.NewCsv(*filename)
	fmt.Printf("filename in Csv struct: %v\n", c.GetFilename())

	data := c.Read()
	for _, line := range data {
		fmt.Println(line)
	}
}
