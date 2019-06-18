package calculatorservice_test

import (
	"fmt"
	"testing"

	"github.com/kasch22/basiccalculator/calculatorservice"
	testUtil "github.com/kasch22/basiccalculator/testing"
)

type testSumString struct {
	sum      string
	expected float64
}

func TestItCalculatesCorrectValueStringSums(t *testing.T) {
	tests := []testSumString{
		{"2 + 2", 4.00},
		{"2 * 2 + 7", 11.00},
		{"2 * 2 + 7 + 10", 21.00},
		{"10 / 4", 2.5},
		{"10 + 4 / 2", 12.0},
		{"10 + 4 / 2 * 4 - 7", 11.0},
		{"10 * ( 7 + 3 )", 100.0},
		{"(10 +2) * ( 7 + 3 )", 120.0},
		{"((4 + 4) + 10 + 2) * ( 7 + 3 )", 200.0},
		{"(((4 + 4) + 10 + 2) - 2) * ( 7 + 3 )", 180.0},
		{"(((4 + 4) + 10 + 2) - 2) * (((7 + 3 ) * 3) / 5 )", 108.0},
		{"(((4 + 4) + 10 + 2) - 2) * (((7 + 3 ) * 3) / 5 ) * 2", 216.0},
		{"4 + 6 + 7", 17.0},
		{"4 + 6 + 7 - 2", 15.0},
		{"(4 * 2) + 7 - 5 * 1", 10.0},
		{"(4 * 2) + (7 - 5) + (5 + 1)", 16.0},
		{"(4 * 2) +(7- (5 + 0)) + (5 + (1 + 1 - (3 - 2)))", 16.0},
		{"(4*2)+(7-(5+0))+(5+(1+1-(3- 2)))", 16.0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Running test %v", test), assertSumHandler(test))
	}
}

func assertSumHandler(test testSumString) func(t *testing.T) {
	return func(t *testing.T) {
		result, err := calculatorservice.CalculateString(test.sum)

		if err != nil {
			t.Fatal(err)
		}

		assertFloatEquals := testUtil.BakeFloat64Equals(t)
		assertFloatEquals("Result", result, test.expected)
	}
}
