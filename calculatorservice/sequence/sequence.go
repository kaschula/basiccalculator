package sequence

import (
	"errors"
	"strings"
)

type Sequence []string

func ToSequence(segments ...string) Sequence {
	s := Sequence{}
	for _, segment := range segments {
		s = s.Append(segment)
	}

	return s
}

func (s Sequence) Append(segments ...string) Sequence {
	return append(s, segments...)
}

func (s Sequence) Split(position int) (Sequence, Sequence) {
	return s[:position], s[position+1:]
}

func (s Sequence) Replace(from, to int, with string) Sequence {
	return append(
		append(s[:from], with), s[to+1:]...,
	)
}

func (s Sequence) ReturnInner(from, to int) (Sequence, error) {
	sequenceLength := len(s)

	if from >= sequenceLength || to >= sequenceLength {
		return nil, errors.New("Bracket Indexes are larger than sequence")
	}

	return s[from+1 : to], nil
}

func (s Sequence) Previous(cutOff int) Sequence {
	return s[:cutOff]
}

func (s Sequence) Remaining(cutOff int) Sequence {
	return s[cutOff+1:]
}

func (s Sequence) Join(after Sequence, with string) Sequence {
	before := append(s, with)

	return append(before, after...)
}

func (s Sequence) IsFirstOrLast(position int) bool {
	return position == 0 || position == (len(s)-1)
}

func (s Sequence) IsSingleValue() bool {
	return len(s) == 1
}

func (s Sequence) IsSimpleSequence() bool {
	return len(s) == 3
}

func BracketLengthsMatch(sequence []string) bool {
	var openList []int
	var closeList []int

	for position, segment := range sequence {
		if strings.Contains("(", segment) {
			openList = append(openList, position)
		}

		if strings.Contains(")", segment) {
			closeList = append(closeList, position)
		}
	}

	return len(openList) == len(closeList)
}

// func ContainsInvalidCharactersOnly(segments []string) bool {
// 	// TODO: need regex
// 	return true
// }
