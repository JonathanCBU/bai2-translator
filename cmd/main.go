package main

import (
	"fmt"
	"io"
	"os"

	"github.com/JonathanCBU/bai2-translator/internal/parser"
	"github.com/joho/godotenv"
)

const (
	TestFilePathName string = "TEST_FILE_PATH"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	file, isSet := os.LookupEnv(TestFilePathName)
	if !isSet {
		fmt.Println("No test file path set")
		return
	}
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
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println(err)
		return
	}

	bai2, err := parser.NewBai2FromReader(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = bai2.Clean()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(bai2.Raw))
	fmt.Println(len(bai2.Data))
}
