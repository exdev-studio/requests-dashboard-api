package model_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/exdev-studio/requests-dashboard-api/internal/model"
)

func TestRequest_Validate(t *testing.T) {
	testCases := []struct {
		name  string
		valid bool
		r     func() *model.Request
		msg   string
	}{
		{
			name:  "valid",
			valid: true,
			r: func() *model.Request {
				return model.TestRequest()
			},
		},
		{
			name:  "empty type",
			valid: false,
			r: func() *model.Request {
				tr := model.TestRequest()
				tr.Type = ""

				return tr
			},
			msg: model.ErrReqTypeFormat,
		},
		{
			name:  "invalid type",
			valid: false,
			r: func() *model.Request {
				tr := model.TestRequest()
				tr.Type = "invalid"

				return tr
			},
			msg: model.ErrReqTypeFormat,
		},
		{
			name:  "empty fields",
			valid: false,
			r: func() *model.Request {
				tr := model.TestRequest()
				tr.Fields = []model.Field{}

				return tr
			},
			msg: model.ErrReqFieldsFormat,
		},
		{
			name:  "fields not set",
			valid: false,
			r: func() *model.Request {
				tr := model.TestRequest()
				tr.Fields = nil

				return tr
			},
			msg: model.ErrReqFieldsFormat,
		},
		{
			name:  "non uniq name in fields",
			valid: false,
			r: func() *model.Request {
				tr := model.TestRequest()
				tf := model.TestField()
				tr.Fields = []model.Field{*tf, *tf}

				return tr
			},
			msg: model.ErrReqFieldsFormat,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.valid {
				assert.NoError(t, tc.r().Validate())
			} else {
				assert.EqualError(t, tc.r().Validate(), errors.New(tc.msg).Error())
			}
		})
	}
}
