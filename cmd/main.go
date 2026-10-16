package main

import (
	"fmt"

	"github.com/JonathanCBU/bai2-translator/internal/parser"
)

func main() {
	s := "01,WPACAU2DK,BHOL,190321,1631,01,080,,2/"
	section, err := parser.NewSection(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(section.Type)

	idx, err := section.GetFieldIndex(parser.RecordCode)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(idx)
}
