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

	txns, err := bai2.GetSectionsByType(parser.T, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, field := range txns[0].Fields {
		fmt.Println(field)
	}
}
