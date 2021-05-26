package main

import (
	"flag"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Zipcode Converter: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	filename := flag.String("filename", "zipcode.csv", "CSV file name of Zipcode.")
	flag.Parse()
	fmt.Printf("file name: %v\n", *filename)
}
