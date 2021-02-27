package model

func TestField() *Field {
	return &Field{
		Type:  FieldString,
		Name:  "test-field",
		Value: "test-value",
	}
}

func TestRequest() *Request {
	f := TestField()
	r := &Request{
		ID:     1,
		Type:   ReqVkCredentials,
		Fields: []Field{*f},
	}

	return r
}
