package parser

import "fmt"

// AccountHeader marks the beginning of an account and provides account number,
// type, currency code, funds type, and item count
type AccountHeader struct {
	Metadata Section
}

func NewAccountHeader(id SectionIdentifier, r string) (*AccountHeader, error) {
	if id != AH {
		return nil, fmt.Errorf(MismatchIdentifierError)
	}
	return &AccountHeader{
		Metadata: *NewSection(r, id, true),
	}, nil
}

// AccountTrailer marks the end of an account and provides account control
// total and number of account records
type AccountTrailer struct {
	Metadata Section
}

func NewAccountTrailer(id SectionIdentifier, r string) (*AccountTrailer, error) {
	if id != AT {
		return nil, fmt.Errorf(MismatchIdentifierError)
	}
	return &AccountTrailer{
		Metadata: *NewSection(r, id, true),
	}, nil
}

// Transaction provides transaction details like type code, amount, funds type,
// westpac ref number, and customer ref number
type Transaction struct {
	Metadata Section
}

func NewTransaction(id SectionIdentifier, r string) (*Transaction, error) {
	if id != T {
		return nil, fmt.Errorf(MismatchIdentifierError)
	}
	return &Transaction{
		Metadata: *NewSection(r, id, false),
	}, nil
}

// Continuation serves to extend Transaction and AccountHeader records to
// multiple lines
type Continuation struct {
	Metadata Section
}

func NewContinuation(id SectionIdentifier, r string) (*Continuation, error) {
	if id != C {
		return nil, fmt.Errorf(MismatchIdentifierError)
	}
	return &Continuation{
		Metadata: *NewSection(r, id, false),
	}, nil
}
