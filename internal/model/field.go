package model

import (
	"errors"
)

const (
	FieldString        = "string"
	ErrFieldTypeFormat = "wrong field type format"
	ErrFieldNameFormat = "wrong field name format"
)

type Field struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (f *Field) Validate() error {
	if f.Type != FieldString {
		return errors.New(ErrFieldTypeFormat)
	}

	if f.Name == "" {
		return errors.New(ErrFieldNameFormat)
	}

	return nil
}
