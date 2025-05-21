package main

import (
	"fmt"
	"github.com/fumiama/go-docx"
	"os"
)

func main() {
	// readFile, err := os.Open("work/sample.docx")
	readFile, err := os.Open("work/sample2.docx")
	if err != nil {
		panic(err)
	}
	fileinfo, err := readFile.Stat()
	if err != nil {
		panic(err)
	}
	size := fileinfo.Size()
	doc, err := docx.Parse(readFile, size)
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain text:")

	for _, it := range doc.Document.Body.Items {
		switch it.(type) {
		case *docx.Paragraph: // printable
			fmt.Println(it)
		}
	}
}
