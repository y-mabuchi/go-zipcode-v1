package main

import (
	"log"
)

func init() {
	log.SetPrefix("Zipcode Converter: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	log.Println("log test")
}
