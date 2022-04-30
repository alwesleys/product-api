package data

import "testing"

func TestValidation(t *testing.T) {
	p := Product{
		Name:  "WesTea",
		Price: 123,
		SKU:   "abs-asd-qwdqwdqwd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
