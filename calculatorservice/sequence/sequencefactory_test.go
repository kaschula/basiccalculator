package sequence_test

import (
	"fmt"
	"testing"

	seq "github.com/kasch22/basiccalculator/calculatorservice/sequence"
	testUtil "github.com/kasch22/basiccalculator/testing"
)

func TestItReturnsASequenceFromASumString(t *testing.T) {
	sum := "2 + 2"
	s, err := seq.CreateSequence(sum)

	if err != nil {
		t.Fatal("Unexpected Error: ", err)
	}

	assertInt := testUtil.BakeIntEquals(t)
	assertEquals := testUtil.BakeStringEquals(t)

	assertInt("Checking Length", 3, (len(s)))
	assertEquals("Checking index 0", "2", s[0])
	assertEquals("Checking index 1", "+", s[1])
	assertEquals("Checking index 2", "2", s[2])
}

func TestItReturnsASequenceFromASumStringWithTwoDigitNumber(t *testing.T) {
	sum := "12 +2"
	s, err := seq.CreateSequence(sum)

	if err != nil {
		t.Fatal("Unexpected Error: ", err)
	}

	assertEquals := testUtil.BakeStringEquals(t)

	assertEquals("Checking index 0", "12", s[0])
	assertEquals("Checking index 2", "+", s[1])
	assertEquals("Checking index 3", "2", s[2])
}

func TestItReturnsASequenceFromASumStringWithThreeDigitNumber(t *testing.T) {
	sum := "123 + 233"
	s, err := seq.CreateSequence(sum)

	if err != nil {
		t.Fatal("Unexpected Error: ", err)
	}

	fmt.Println(s)

	assertEquals := testUtil.BakeStringEquals(t)

	assertEquals("Checking index 0", "123", s[0])
	assertEquals("Checking index 1", "+", s[1])
	assertEquals("Checking index 2", "233", s[2])
}

func TestItReturnsASequenceFromASumWithBrackets(t *testing.T) {
	sum := "12 + 2 * (4 +381)"
	s, err := seq.CreateSequence(sum)

	if err != nil {
		t.Fatal("Unexpected Error: ", err)
	}
	assertEquals := testUtil.BakeStringEquals(t)

	assertEquals("Checking index 0", "12", s[0])
	assertEquals("Checking index 1", "+", s[1])
	assertEquals("Checking index 2", "2", s[2])
	assertEquals("Checking index 3", "*", s[3])
	assertEquals("Checking index 4", "(", s[4])
	assertEquals("Checking index 5", "4", s[5])
	assertEquals("Checking index 6", "+", s[6])
	assertEquals("Checking index 7", "381", s[7])
	assertEquals("Checking index 8", ")", s[8])
}

func TestItReturnsASequenceWithANegativeNumber(t *testing.T) {
	sum := "123 + -233"
	s, err := seq.CreateSequence(sum)

	if err != nil {
		t.Fatal("Unexpected Error: ", err)
	}

	assertEquals := testUtil.BakeStringEquals(t)

	assertEquals("Checking index 0", "123", s[0])
	assertEquals("Checking index 1", "+", s[1])
	assertEquals("Checking index 2", "-233", s[2])
}

func TestItReturnsASequenceStartingWithANegativeNumber(t *testing.T) {
	sum := "-123 +-233"
	s, err := seq.CreateSequence(sum)

	if err != nil {
		t.Fatal("Unexpected Error: ", err)
	}

	assertEquals := testUtil.BakeStringEquals(t)

	assertEquals("Checking index 0", "-123", s[0])
	assertEquals("Checking index 1", "+", s[1])
	assertEquals("Checking index 2", "-233", s[2])
}

func TestItFailsForInvalidSequenceOfEvenLength(t *testing.T) {
	sum := "-123 +-233)"
	_, err := seq.CreateSequence(sum)

	if err == nil {
		t.Fatal("Expecting and Error")
	}

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("Should return a length err", "Error: Sequence Length must be an off number", err.Error())
}

func TestItFailsForSequencesWithUnevenBracketsCounts(t *testing.T) {
	sum := "-123 +(233))"
	_, err := seq.CreateSequence(sum)

	if err == nil {
		t.Fatal("Expecting and Error")
	}

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("Should return a length err", "Error: Sequence Length must be an off number", err.Error())
}
