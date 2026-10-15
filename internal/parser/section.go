package parser

type SectionIdentifier string

const (
	FH SectionIdentifier = "01"
	GH SectionIdentifier = "02"
	AH SectionIdentifier = "03"
	T  SectionIdentifier = "16"
	AT SectionIdentifier = "49"
	GT SectionIdentifier = "98"
	FT SectionIdentifier = "99"
	C  SectionIdentifier = "88"
)

type Section struct {
	Raw       string
	Type      SectionIdentifier
	Mandatory bool
}

func NewSection(r string, t SectionIdentifier, m bool) *Section {
	return &Section{
		Raw:       r,
		Type:      t,
		Mandatory: m,
	}
}

type ISection interface {
	GetFieldIndex(t FieldType) (int, error)
	GetFieldValueByType(t FieldType) (string, error)
	GetFieldValueByIndex(i int) (string, error)
}
