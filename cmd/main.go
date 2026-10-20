package main

import (
	"fmt"
	"os"

	"github.com/JonathanCBU/bai2-translator/internal/parser"
)

func main() {
	file := "./tmp/bai2-npp-sample-file.txt"
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	sections, err := parser.ScanFileToSections(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range sections {
		fmt.Println(sections[i].Type)
	}
}
