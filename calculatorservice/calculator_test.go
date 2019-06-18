package calculatorservice_test

import (
	"fmt"
	"testing"

	"github.com/kasch22/basiccalculator/calculatorservice"
	seq "github.com/kasch22/basiccalculator/calculatorservice/sequence"
	testUtil "github.com/kasch22/basiccalculator/testing"
)

// Should be lower case
type TestSum struct {
	sum      seq.Sequence
	expected float64
}

func TestItCalculatesCorrectValueForSimpleSums(t *testing.T) {
	tests := []TestSum{
		{newSeq("2", "+", "2"), 4.00},
		{newSeq("2", "*", "2", "+", "7"), 11.00},
		{newSeq("2", "*", "2", "+", "7", "+", "10"), 21.00},
		{newSeq("10", "/", "4"), 2.5},
		{newSeq("10", "+", "4", "/", "2"), 12.0},
		{newSeq("10", "+", "4", "/", "2", "*", "4", "-", "7"), 11.0},
		{newSeq("10", "*", "(", "7", "+", "3", ")"), 100.0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Running test %v", test), assertResultHandler(test))
	}
}

func TestItCalculatesBracketSumsCorrectValue(t *testing.T) {
	tests := []TestSum{
		{newSeq("11", "*", "(", "7", "+", "3", ")"), 110},
		{newSeq("10", "*", "(", "7", "+", "(", "3", "+", "4", ")", ")"), 140},
		{newSeq("10", "*", "(", "(", "4", "+", "4", ")", "+", "8", ")"), 160},
		{newSeq("11", "+", "2", "*", "(", "3", "+", "7", ")"), 31},
		{newSeq("(", "3", "*", "(", "4", "+", "5", ")", ")"), 27},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Running test %v", test), assertResultHandler(test))
	}
}

func TestItCalculatesMultipleBracketSetSums(t *testing.T) {
	tests := []TestSum{
		{newSeq("(", "3", "*", "(", "4", "+", "5", ")", ")", "+", "(", "12", "+", "(", "20", "/", "2", ")", ")"), 49},
		{newSeq("(", "3", "+", "4", ")", "+", "(", "12", "+", "1", ")"), 20},
		{newSeq("(", "3", "+", "4", ")", "+", "(", "1", "+", "(", "4", "+", "8", ")", ")"), 20},
		{newSeq("(", "3", "+", "(", "20", "/", "5", ")", ")", "+", "(", "1", "+", "(", "4", "+", "8", ")", ")"), 20},
		{newSeq("(", "2", "+", "3", ")", "+", "(", "3", "+", "4", ")", "+", "(", "4", "+", "5", ")"), 21},
		{newSeq(
			"(", "2", "+", "(", "3", "+", "7", ")", ")",
			"+",
			"(", "3", "+", "4", ")",
			"+",
			"(", "4", "+", "5", ")",
		),
			28,
		},
		{newSeq(
			"(", "2", "+", "(", "3", "+", "7", ")", ")",
			"+",
			"(", "3", "+", "(", "4", "-", "1", ")", ")",
			"+",
			"(", "4", "+", "5", ")",
		),
			27,
		},
		{newSeq(
			"(", "2", "+", "(", "3", "+", "7", ")", ")",
			"+",
			"(", "3", "+", "(", "4", "-", "1", ")", ")",
			"+",
			"(", "4", "+", "(", "5", "-", "4", ")", ")",
		),
			23,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Running test %v", test), assertResultHandler(test))
	}
}

func assertResultHandler(test TestSum) func(t *testing.T) {
	return func(t *testing.T) {
		result, err := calculatorservice.Calculate(test.sum)

		if err != nil {
			t.Fatal(err)
		}

		assertFloatEquals := testUtil.BakeFloat64Equals(t)
		assertFloatEquals("Result", result, test.expected)
	}
}

func newSeq(values ...string) seq.Sequence {
	return seq.ToSequence(values...)
}
