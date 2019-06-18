package bracketset

import (
	"errors"
	"strings"

	seq "github.com/kasch22/basiccalculator/calculatorservice/sequence"
)

type BracketSet struct {
	open  int
	close int
}

func (bs BracketSet) IsEmpty() bool {
	return bs.open == bs.close
}

func (bs BracketSet) Open() int {
	return bs.open
}

func (bs BracketSet) Close() int {
	return bs.close
}

func InnerMostBracket(sequence seq.Sequence) (BracketSet, error) {
	var openList []int
	var closeList []int

	for position, segment := range sequence {
		if strings.Contains("(", segment) {
			openList = append(openList, position)
		}

		if strings.Contains(")", segment) {
			closeList = append([]int{position}, closeList...)
		}
	}

	if len(openList) != len(closeList) {
		return BracketSet{0, 0}, errors.New("Error: open and close brackets should be equal")
	}

	if len(openList) == 0 {
		return BracketSet{0, 0}, nil
	}

	lastItemIndex := len(closeList) - 1

	return BracketSet{openList[lastItemIndex], closeList[lastItemIndex]}, nil
}
