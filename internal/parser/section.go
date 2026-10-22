package parser

import (
	"fmt"
	"strings"
)

// SectionIdentifier is the string at the front of a line in a BAI2 file.
// It tells us what fields to expect from that row.
type SectionIdentifier string

const (
	FH SectionIdentifier = "01" // file header
	GH SectionIdentifier = "02" // group header
	AH SectionIdentifier = "03" // account header
	T  SectionIdentifier = "16" // transaction
	AT SectionIdentifier = "49" // account trailer
	GT SectionIdentifier = "98" // group trailer
	FT SectionIdentifier = "99" // file trailer
	C  SectionIdentifier = "88" // continuation (of AH or T)
)

var validSectionIdentifiers = [8]SectionIdentifier{
	FH,
	GH,
	AH,
	T,
	C,
	AT,
	GT,
	FT,
}

var mandatoryMap = map[SectionIdentifier]bool{
	FH: true,
	GH: true,
	AH: true,
	T:  false,
	C:  false,
	AT: true,
	GT: true,
	FT: true,
}

// Section represents one row of a BAI2 file
type Section struct {
	Raw       string
	Type      SectionIdentifier
	Mandatory bool
	Fields    []string
	NumFields int
}

func NewSection(r string) (*Section, error) {
	trimmed := strings.TrimRight(r, "/")
	f := strings.Split(trimmed, ",")

	t, err := GetSectionIdentifier(r)
	if err != nil {
		return nil, err
	}

	m, err := IsSectionMandatory(t)
	if err != nil {
		return nil, err
	}

	return &Section{
		Raw:       r,
		Type:      t,
		Mandatory: m,
		Fields:    f,
	}, nil
}

type ISection interface {
	GetFieldIndex(t FieldType) (int, error)
	GetFieldValueByType(t FieldType) (string, error)
	GetFieldValueByIndex(i int) (string, error)
}

func (s *Section) GetFieldIndex(t FieldType) (int, error) {
	f, err := getSectionFields(s.Type)
	if err != nil {
		return -1, err
	}
	return getFieldIndexOrErr(f, t)
}

func (s *Section) GetFieldValueByType(t FieldType) (string, error) {
	idx, err := s.GetFieldIndex(t)
	if err != nil {
		return "", err
	}
	if idx >= len(s.Fields) {
		return "", fmt.Errorf(FieldIndexOutOfRange)
	}
	return s.Fields[idx], nil
}

func (s *Section) GetFieldValueByIndex(i int) (string, error) {
	if i >= len(s.Fields) {
		return "", fmt.Errorf(FieldIndexOutOfRange)
	}
	return s.Fields[i], nil
}

func (s *Section) Merge(ext *Section) (bool, error) {
	if ext.Type != C {
		return false, fmt.Errorf(CanOnlyMergeContinuations)
	}
	initialSize := len(s.Fields)
	s.Fields = append(s.Fields, ext.Fields[1:]...)
	return initialSize != len(s.Fields), nil
}

func IsSectionMandatory(t SectionIdentifier) (bool, error) {
	m, exists := mandatoryMap[t]
	if !exists {
		return false, fmt.Errorf(UnknownSectionIdentifier)
	}
	return m, nil
}

func GetSectionIdentifier(s string) (SectionIdentifier, error) {
	trimmed := strings.SplitN(s, ",", 2)

	for _, sid := range validSectionIdentifiers {
		if string(sid) == trimmed[0] {
			return sid, nil
		}
	}
	return "", fmt.Errorf(UnknownSectionIdentifier)
}
