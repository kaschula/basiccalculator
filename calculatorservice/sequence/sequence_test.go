package sequence_test

import (
	"testing"

	"github.com/kasch22/basiccalculator/calculatorservice/sequence"
	testUtil "github.com/kasch22/basiccalculator/testing"
)

func TestItCreatesANewSequenceWithOneItem(t *testing.T) {
	sequence := sequence.ToSequence("1")

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("Sequence index 0", "1", sequence[0])
}

func TestItCreatesANewSequenceWithMultipleItems(t *testing.T) {
	sequence := sequence.ToSequence("1", "+", "2")

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("Sequence index 0", "1", sequence[0])
	assertEquals("Sequence index 1", "+", sequence[1])
	assertEquals("Sequence index 1", "2", sequence[2])
}

func TestItCanAppendAnItem(t *testing.T) {
	sequence := sequence.ToSequence("1")
	sequence = sequence.Append("2")

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("Sequence index 0", "1", sequence[0])
	assertEquals("Sequence index 1", "2", sequence[1])
}

func TestItCanAppendMultiple(t *testing.T) {
	sequence := sequence.ToSequence("1")
	sequence = sequence.Append("2", "3")

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("Sequence index 0", "1", sequence[0])
	assertEquals("Sequence index 1", "2", sequence[1])
	assertEquals("Sequence index 2", "3", sequence[2])
}

func TestItCanSplitByPosition(t *testing.T) {
	sequence := sequence.ToSequence("1", "+", "3")
	before, after := sequence.Split(1)

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("Before index 0", "1", before[0])
	assertEquals("After index 0", "3", after[0])
}

func TestItCanSplitByPositionWithLargeSum(t *testing.T) {
	sequence := sequence.ToSequence("1", "+", "3", "*", "40", "/", "6")
	before, after := sequence.Split(3)

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("Before index 0", "1", before[0])
	assertEquals("Before index 1", "+", before[1])
	assertEquals("Before index 2", "3", before[2])
	assertEquals("After index 0", "40", after[0])
	assertEquals("After index 1", "/", after[1])
	assertEquals("After index 2", "6", after[2])
}

func TestItInsertANewValueBetweenTwoPoints(t *testing.T) {
	sequence := sequence.ToSequence("1", "+", "3", "*", "4", "/", "6")
	replaced := sequence.Replace(2, 4, "12")

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("replaced index 0", "1", replaced[0])
	assertEquals("replaced index 1", "+", replaced[1])
	assertEquals("replaced index 2", "12", replaced[2])
	assertEquals("replaced index 3", "/", replaced[3])
	assertEquals("replaced index 4", "6", replaced[4])
}

func TestItReturnsNewSequenceBetweenTwoPositons(t *testing.T) {
	sequence := sequence.ToSequence("1", "+", "(", "4", "/", "6", ")")
	innerSequence, err := sequence.ReturnInner(2, 6)

	if err != nil {
		t.Fatal("Error: ", err)
	}

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("replaced index 0", "4", innerSequence[0])
	assertEquals("replaced index 1", "/", innerSequence[1])
	assertEquals("replaced index 2", "6", innerSequence[2])
}

// func TestItCanGetSequenceBeforeAPosition
func TestItCanGetSequenceBeforeAPosition(t *testing.T) {
	sequence := sequence.ToSequence("1", "+", "(", "4", "/", "6", ")")

	previous := sequence.Previous(2)

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("replaced index 0", "1", previous[0])
	assertEquals("replaced index 1", "+", previous[1])
}

func TestItCanGetTheRemainingSequenceAPosition(t *testing.T) {
	sequence := sequence.ToSequence("(", "4", "/", "6", ")", "+", "2")

	remaining := sequence.Remaining(4)

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("replaced index 0", "+", remaining[0])
	assertEquals("replaced index 1", "2", remaining[1])
}

// test it can merge to sequences with a value in between

func TestItCanJoinTwoSequencesWithAValueInbetween(t *testing.T) {
	before := sequence.ToSequence("4", "/", "6")
	after := sequence.ToSequence("5", "+", "7")
	joined := before.Join(after, "+")

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("replaced index 0", "4", joined[0])
	assertEquals("replaced index 1", "/", joined[1])
	assertEquals("replaced index 2", "6", joined[2])
	assertEquals("replaced index 3", "+", joined[3])
	assertEquals("replaced index 4", "5", joined[4])
	assertEquals("replaced index 5", "+", joined[5])
	assertEquals("replaced index 6", "7", joined[6])
}

func TestItCanTellIfPositionGivenIsTheFirstOrLastSliceElement(t *testing.T) {
	s := sequence.ToSequence("4", "/", "6")

	assertEquals := testUtil.BakeBoolEquals(t)
	assertEquals("CheckingFirst", true, s.IsFirstOrLast(0))
	assertEquals("Checking Middle", false, s.IsFirstOrLast(1))
	assertEquals("CheckingLast", true, s.IsFirstOrLast(2))
}

func TestASequenceOfOneValueResolvesTrue(t *testing.T) {
	singleValue := sequence.ToSequence("4")
	simpleSequence := sequence.ToSequence("4", "/", "6")

	assertEquals := testUtil.BakeBoolEquals(t)
	assertEquals("Checking Single Value Sequence", true, singleValue.IsSingleValue())
	assertEquals("Checking a Simple Sequence", false, simpleSequence.IsSingleValue())
}
func TestASimpleSequenceOfOneValueResolvesTrue(t *testing.T) {
	singleValue := sequence.ToSequence("4")
	simpleSequence := sequence.ToSequence("4", "/", "6")
	complexSequence := sequence.ToSequence("4", "/", "6", "+", "2")

	assertEquals := testUtil.BakeBoolEquals(t)
	assertEquals("Checking Single Value Sequence", false, singleValue.IsSimpleSequence())
	assertEquals("Checking a Simple Sequence", true, simpleSequence.IsSimpleSequence())
	assertEquals("Checking Checking a complex sequence", false, complexSequence.IsSimpleSequence())
}
