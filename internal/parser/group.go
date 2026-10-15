package parser

import "fmt"

// GroupHeader marks the beginning of a group. and identifies the group of
// accounts sent in the file that are from the same sender along with date of
// origin and currency code for sender's country
type GroupHeader struct {
	Metadata Section
}

func NewGroupHeader(id SectionIdentifier, r string) (*GroupHeader, error) {
	if id != GH {
		return nil, fmt.Errorf(MismatchIdentifierError)
	}
	return &GroupHeader{
		Metadata: *NewSection(r, id, true),
	}, nil
}

// GroupTrailer marks a group end and contains group control total, account
// num and record num of the group
type GroupTrailer struct {
	Metadata Section
}

func NewGroupTrailer(id SectionIdentifier, r string) (*GroupTrailer, error) {
	if id != GT {
		return nil, fmt.Errorf(MismatchIdentifierError)
	}
	return &GroupTrailer{
		Metadata: *NewSection(r, id, true),
	}, nil
}
