package parser

// FieldType is an enum of unique BAI2 fields. Some types cover multiple
// SectionTypes
type FieldType int

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
