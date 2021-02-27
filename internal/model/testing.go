package model

func TestField() *Field {
	return &Field{
		Type:  FieldString,
		Name:  "test-field",
		Value: "test-value",
	}
}
