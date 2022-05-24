package calculator_test

import (
	"github.com/chkda/go-testing/calculator"
	"testing"
)

type TestCase struct {
	name     string
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsArmstrong(t *testing.T) {
	t.Run("test for all 3 digit armstrong numbers", func(t *testing.T) {
		tests := []TestCase{
			{name: "Testing value for : 153", value: 153, expected: true},
			{name: "Testing value for : 370", value: 370, expected: true},
			{name: "Testing value for : 371", value: 371, expected: true},
			{name: "Testing value for : 407", value: 407, expected: true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual := calculator.CalculateIsArmstrong(test.value)
				if actual != test.expected {
					t.Fail()
				}
			})
		}
	})
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {
	t.Run("Should fail for case 350", func(t *testing.T) {
		test := TestCase{
			value:    350,
			expected: false,
		}

		actual := calculator.CalculateIsArmstrong(test.value)
		if actual != test.expected {
			t.Fail()
		}
	})

	t.Run("Should fail for case 200", func(t *testing.T) {
		test := TestCase{
			value:    200,
			expected: false,
		}

		actual := calculator.CalculateIsArmstrong(test.value)
		if actual != test.expected {
			t.Fail()
		}
	})
}
