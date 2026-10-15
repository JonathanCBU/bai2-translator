package main

import (
	"fmt"

	"github.com/JonathanCBU/bai2-translator/internal/parser"
)

func main() {
	s := "01,WPACAU2DK,BHOL,190321,1631,01,080,,2/"
	f, err := parser.NewFileHeader(parser.FH, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	i, err := f.GetFieldIndex(parser.SenderID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)

	fmt.Println(f.GetFieldValueByType(parser.SenderID))
	fmt.Println(f.GetFieldValueByIndex(4))
}
