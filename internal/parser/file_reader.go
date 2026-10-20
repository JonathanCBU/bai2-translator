package parser

import (
	"bufio"
	"fmt"
	"io"
)

func ScanFileToSections(f io.Reader) ([]*Section, error) {
	var sections []*Section
	lineNum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sec, err := NewSection(line)
		if err != nil {
			fmt.Printf("failed to parse line %d\n%s\n\n", lineNum, line)
		} else {
			sections = append(sections, sec)
		}
		lineNum++
	}
	return sections, scanner.Err()
}
