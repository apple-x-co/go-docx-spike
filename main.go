package main

import (
	"baliance.com/gooxml/document"
	"log"
	"os"
)

func main() {
	doc, err := document.Open("work/sample2.docx")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("work/gooxml/document/output.docx")
	if err != nil {
		panic(err)
	}
	err = doc.Save(f)
	if err != nil {
		panic(err)
	}
}
