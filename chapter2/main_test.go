package main

import (
	. "chapter2/helpers"
	"errors"
	"testing"
)

var tests = []struct {
	name   string
	a      float32
	b      float32
	result float32
	err    error
}{
	{"divide 2 by 1", 2, 1, 2, nil},
	{"divide 2 by 0", 2, 0, 0, errors.New("cannot divide by zero")},
}

func TestDivision(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := CalculateDivide(test.a, test.b)
			if result != test.result {
				t.Errorf("Expected result to equal %f, but got %f", test.result, result)
			}
			if err != nil {
				if err.Error() != test.err.Error() {
					t.Errorf("Expected error to equal %s, but got %s", test.err.Error(), err.Error())
				}
			}
		})
	}
}

func TestCalculateDivideOk(t *testing.T) {
	var result, err = CalculateDivide(2, 1)

	if err != nil {
		t.Error("Expected err to be nil")
	}

	if result != 2 {
		t.Error("Expected result to equal 2")
	}
}

func TestCalculateDivideError(t *testing.T) {
	var _, err = CalculateDivide(2, 0)

	if err == nil {
		t.Error("Expected error not to be nil")
	}
}
