package bracketset_test

import (
	"testing"

	"github.com/kasch22/basiccalculator/calculatorservice/bracketset"
	"github.com/kasch22/basiccalculator/calculatorservice/sequence"
	testUtil "github.com/kasch22/basiccalculator/testing"
)

func TestItReturnsABracketSet(t *testing.T) {
	s := sequence.ToSequence("(", "2", "+", "3", ")")

	bracketSet, err := bracketset.InnerMostBracket(s)

	if err != nil {
		t.Fatal(err)
	}

	assertEquals := testUtil.BakeIntEquals(t)
	assertEquals("open position", bracketSet.Open(), 0)
	assertEquals("close position", bracketSet.Close(), 4)
}

func TestItReturnsAnEmptyBracketSetWhenSequenceHasNoBracket(t *testing.T) {
	s := sequence.ToSequence("2", "+", "3")

	bracketSet, err := bracketset.InnerMostBracket(s)

	if err != nil {
		t.Fatal(err)
	}

	assertEquals := testUtil.BakeBoolEquals(t)
	assertEquals("is Empty", bracketSet.IsEmpty(), true)
}

func TestTheSequenceInsideABracketCanBeResolved(t *testing.T) {
	data := sequence.ToSequence("(", "2", "+", "3", ")")

	bracket, err := bracketset.InnerMostBracket(data)

	if err != nil {
		t.Fatal(err, err)
	}

	sequence, err := data.ReturnInner(bracket.Open(), bracket.Close())

	if err != nil {
		t.Fatal(err)
	}

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("1st position", sequence[0], "2")
	assertEquals("2nd position", sequence[1], "+")
	assertEquals("3rd position", sequence[2], "3")
}

func TestTheInnerMostBracketIsReturnedFirst(t *testing.T) {
	data := sequence.ToSequence("(", "2", "+", "(", "4", "*", "3", ")", ")")

	bracket, err := bracketset.InnerMostBracket(data)
	if err != nil {
		t.Fatal(err)
	}

	sequence, err := data.ReturnInner(bracket.Open(), bracket.Close())
	if err != nil {
		t.Fatal("Error creating sequence", err)
	}

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("1st position", sequence[0], "4")
	assertEquals("2nd position", sequence[1], "*")
	assertEquals("3rd position", sequence[2], "3")
}

func TestTheInnerMostBracketIsReturnedFirstInAThreeBracketSet(t *testing.T) {
	data := sequence.ToSequence("(", "2", "+", "(", "4", "*", "(", "3", "/", "5", ")", ")", ")")

	bracket, err := bracketset.InnerMostBracket(data)
	if err != nil {
		t.Fatal(err, err)
	}

	sequence, err := data.ReturnInner(bracket.Open(), bracket.Close())
	if err != nil {
		t.Fatal("Error creating sequence", err)
	}

	assertEquals := testUtil.BakeStringEquals(t)
	assertEquals("1st position", sequence[0], "3")
	assertEquals("2nd position", sequence[1], "/")
	assertEquals("3rd position", sequence[2], "5")
}
