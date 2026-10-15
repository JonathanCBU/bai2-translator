package parser

import (
	"fmt"
	"slices"
)

func getFieldIndexOrErr(sectionFields []FieldType, t FieldType) (int, error) {
	idx := slices.Index(sectionFields, t)
	if idx == -1 {
		return -1, fmt.Errorf(FieldNotInSection)
	}
	return idx, nil
}
