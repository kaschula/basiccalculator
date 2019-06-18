package sequence

import (
	"errors"
	"strings"
)

func CreateSequence(sum string) (Sequence, error) {
	return validate(
		addWhiteSpaceBetweenOperators([]byte(sum)),
	)
}

func addWhiteSpaceBetweenOperators(sequence []byte) string {
	lastPosition := len(sequence) - 1
	var withSpace []byte

	for position, char := range sequence {
		withSpace = append(withSpace, char)

		if position == lastPosition {
			continue
		}

		if isNegativeNumber(position, char, sequence) {
			continue
		}

		if char == ' ' {
			continue
		}

		if isOperatorOrBracket(string(char)) && isOperatorOrBracket(nextChar(position, sequence)) {
			withSpace = append(withSpace, ' ')
			continue
		}

		if isNumber(string(char)) && isOperatorOrBracket(nextChar(position, sequence)) {
			withSpace = append(withSpace, ' ')
			continue
		}

		if isOperatorOrBracket(string(char)) && isNumber(nextChar(position, sequence)) {
			withSpace = append(withSpace, ' ')
			continue
		}
	}

	return string(withSpace)
}

func validate(sum string) (Sequence, error) {
	segments := strings.Split(sum, " ")

	if len(segments)%2 == 0 {
		return nil, errors.New("Error: Sequence Length must be an odd number")
	}

	return ToSequence(segments...), nil
}

func isOperatorOrBracket(char string) bool {
	return strings.Contains("+-/*()", char)
}

func isNumber(char string) bool {
	return !isOperatorOrBracket(char) && char != " "
}

func nextChar(currentPosition int, sum []byte) string {
	return string(sum[currentPosition+1])
}

func isNegativeNumber(position int, char byte, sum []byte) bool {
	return char == '-' && isNumber(nextChar(position, sum))
}
