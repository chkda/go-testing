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

func benchmarkCalculateIsArmstrong(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculator.CalculateIsArmstrong(input)
	}
}

func BenchmarkCalculateIsArmstrong370(b *testing.B) {
	benchmarkCalculateIsArmstrong(370, b)
}

func BenchmarkCalculateIsArmstrong371(b *testing.B) {
	benchmarkCalculateIsArmstrong(371, b)
}

func BenchmarkCalculateIsArmstrong0(b *testing.B) {
	benchmarkCalculateIsArmstrong(0, b)
}
