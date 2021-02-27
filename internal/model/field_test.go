package model_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/exdev-studio/requests-dashboard-api/internal/model"
)

func TestField_Validate(t *testing.T) {
	testCases := []struct {
		name  string
		f     func() *model.Field
		valid bool
		msg   string
	}{
		{
			name: "valid",
			f: func() *model.Field {
				return model.TestField()
			},
			valid: true,
		},
		{
			name: "empty type",
			f: func() *model.Field {
				f := model.TestField()
				f.Type = ""

				return f
			},
			valid: false,
			msg:   model.ErrFieldTypeFormat,
		},
		{
			name: "invalid type",
			f: func() *model.Field {
				f := model.TestField()
				f.Type = "invalid"

				return f
			},
			valid: false,
			msg:   model.ErrFieldTypeFormat,
		},
		{
			name: "empty name",
			f: func() *model.Field {
				f := model.TestField()
				f.Name = ""

				return f
			},
			valid: false,
			msg:   model.ErrFieldNameFormat,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.f().Validate()
			if tc.valid {
				assert.NoError(t, tc.f().Validate())
			} else {
				assert.EqualError(t, err, errors.New(tc.msg).Error())
			}
		})
	}
}
