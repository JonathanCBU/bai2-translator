package parser

import (
	"bufio"
	"fmt"
	"io"
)

type Bai2 struct {
	Raw  []*Section
	Data []*Section
}

type IBai2 interface {
	Print(fmt string)
	Clean() error
}

func NewBai2FromReader(f io.Reader) (*Bai2, error) {
	var raw []*Section
	lineNum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sec, err := NewSection(line)
		if err != nil {
			fmt.Printf("failed to parse line %d\n%s\n\n", lineNum, line)
		} else {
			raw = append(raw, sec)
		}
		lineNum++
	}
	return &Bai2{Raw: raw}, scanner.Err()
}

func (b *Bai2) Clean() error {
	var data []*Section
	var lastSection *Section
	for _, s := range b.Raw {
		if s.Type == C {
			lastSection.Merge(s)
		} else {
			lastSection = s
			data = append(data, s)
		}
	}
	b.Data = data
	return nil
}

func (b *Bai2) GetSectionsByType(t SectionIdentifier, clean bool) ([]*Section, error) {
	var results []*Section
	searchData := b.Raw
	if clean {
		searchData = b.Data
	}

	for _, s := range searchData {
		if s.Type == t {
			results = append(results, s)
		}
	}
	return results, nil
}

func (b *Bai2) Print(fmt string)
