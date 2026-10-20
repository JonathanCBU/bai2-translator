package parser

import (
	"fmt"
	"slices"
)
// FieldType is an enum of unique BAI2 fields. Some types cover multiple
// SectionTypes
type FieldType int

// FieldType is an enum of unique BAI2 fields. Some types cover multiple
// SectionTypes type FieldType int
const (
	RecordCode FieldType = iota
	SenderID
	RecID
	CreatedDate
	CreatedTime
	FileID
	RecordLen
	BlockSize
	BaiVersion
	OriginatorID
	GroupStatus
	AsOfDate
	AsOfTime
	CurrencyCode
	AsOfDateMod
	AcctNum
	AcctType
	Amount
	ItemCount
	FundsType
	TxnType
	BankRef
	CustomerRef
	Text
	ControlTotal
	NumberOfGroups
	NumberOfRecords
)

type Field struct {
	Index     int
	Type      FieldType
	Mandatory bool
	Value     string
}

/*
Section field functions
*/
func getFieldIndexOrErr(sectionFields []FieldType, t FieldType) (int, error) {
	idx := slices.Index(sectionFields, t)
	if idx == -1 {
		return -1, fmt.Errorf(FieldNotInSection)
	}
	return idx, nil
}

func getSectionFields(si SectionIdentifier) ([]FieldType, error) {
	switch si {
	case FH:
		return []FieldType{
			RecordCode,
			SenderID,
			RecID,
			CreatedDate,
			CreatedTime,
			FileID,
			RecordLen,
			BlockSize,
			BaiVersion,
		}, nil
	case FT:
		return []FieldType{
			RecordCode,
			ControlTotal,
			NumberOfGroups,
			NumberOfRecords,
		}, nil
	default:
		return nil, fmt.Errorf(UnknownSectionIdentifier)
	}
}
