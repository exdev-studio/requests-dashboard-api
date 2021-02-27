package model

import (
	"errors"
)

const (
	ReqVkCredentials   = "VK_CREDENTIALS"
	ErrReqTypeFormat   = "wrong type format"
	ErrReqFieldsFormat = "wrong fields format"
)

type Request struct {
	ID     int     `json:"id"`
	Type   string  `json:"type"`
	Fields []Field `json:"fields"`
}

func (r *Request) Validate() error {
	if r.Type != ReqVkCredentials {
		return errors.New(ErrReqTypeFormat)
	}

	if r.Fields == nil || len(r.Fields) <= 0 {
		return errors.New(ErrReqFieldsFormat)
	}

	for _, f := range r.Fields {
		if err := f.Validate(); err != nil {
			return err
		}
	}

	if err := r.validateFieldsUniqueness(); err != nil {
		return err
	}

	return nil
}

func (r *Request) validateFieldsUniqueness() error {
	uniqueMap := make(map[string]bool, 0)

	for _, f := range r.Fields {
		_, ok := uniqueMap[f.Name]
		if ok {
			return errors.New(ErrReqFieldsFormat)
		}

		uniqueMap[f.Name] = true
	}

	return nil
}
