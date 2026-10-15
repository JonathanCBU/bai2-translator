package parser

import (
	"fmt"
	"strings"
)

// FileHeader marks the beginning of a file and contains the sender,
// receiver, and structure description
type FileHeader struct {
	Metadata Section
	Fields   []string
}

var FileHeaderFields = []FieldType{
	RecordCode,
	SenderID,
	RecID,
	CreatedDate,
	CreatedTime,
	FileID,
	RecordLen,
	BlockSize,
	BaiVersion,
}

var NumFileHeaderFields = len(FileHeaderFields)

func NewFileHeader(id SectionIdentifier, r string) (*FileHeader, error) {
	if id != FH {
		return nil, fmt.Errorf(MismatchIdentifierError)
	}
	trimmed := strings.TrimRight(r, "/")
	return &FileHeader{
		Metadata: *NewSection(r, id, true),
		Fields:   strings.Split(trimmed, ","),
	}, nil
}

func (h *FileHeader) GetFieldIndex(t FieldType) (int, error) {
	return getFieldIndexOrErr(FileHeaderFields, t)
}

func (h *FileHeader) GetFieldValueByType(t FieldType) (string, error) {
	i, err := h.GetFieldIndex(t)
	if err != nil {
		return "", err
	}
	return h.Fields[i], nil
}

func (h *FileHeader) GetFieldValueByIndex(i int) (string, error) {
	if i < 0 || i >= NumFileHeaderFields {
		return "", fmt.Errorf(FieldNotInSection)
	}
	return h.Fields[i], nil
}

// FileTrailer marks the end of a file and contains the control total, num of
// accounts, and num of records
type FileTrailer struct {
	Metadata Section
}

var FileTrailerFields = []FieldType{
	RecordCode,
	ControlTotal,
	NumberOfGroups,
	NumberOfRecords,
}

var NumFileTrailerFields = len(FileTrailerFields)

func NewFileTrailer(id SectionIdentifier, r string) (*FileTrailer, error) {
	if id != FT {
		return nil, fmt.Errorf(MismatchIdentifierError)
	}
	return &FileTrailer{
		Metadata: *NewSection(r, id, true),
	}, nil
}
