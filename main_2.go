package main

import (
	"github.com/gomutex/godocx"
	"log"
)

func main() {
	document, err := godocx.OpenDocument("work/sample2.docx")
	if err != nil {
		log.Fatal(err)
	}

	err = document.SaveTo("work/gomutex/godocx/output.docx")
	if err != nil {
		log.Fatal(err)
	}
}
