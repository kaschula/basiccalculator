package calculatorservice

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/kasch22/basiccalculator/calculatorservice/bracketset"
	seq "github.com/kasch22/basiccalculator/calculatorservice/sequence"
)

func Calculate(sequence seq.Sequence) (float64, error) {
	if isMultipleBracketSetSum(sequence) {
		return calculateMultipleBracketSequence(sequence)
	}

	return calculateSingleBracketSequence(sequence)
}

func calculateMultipleBracketSequence(sequence seq.Sequence) (float64, error) {
	for position, segment := range sequence {

		if sequence.IsFirstOrLast(position) {
			continue
		}

		if !isValidOperatorBetweenBrackets(segment, sequence[previous(position)], sequence[next(position)]) {
			continue
		}

		firstBracketSet := sequence.Previous(position)
		secondBracketSet := sequence.Remaining(position)

		resultOne, errOne := Calculate(firstBracketSet)
		resultTwo, errTwo := Calculate(secondBracketSet)

		if errOne != nil || errTwo != nil {
			return 0.00, errors.New(fmt.Sprintf("calculation errors: %v %v", errOne, errTwo))
		}

		return calculateSingleBracketSequence(
			seq.ToSequence(fmt.Sprintf("%v", resultOne), segment, fmt.Sprintf("%v", resultTwo)),
		)
	}

	return 0.00, errors.New("Could not calculate")
}

func calculateSingleBracketSequence(sequence seq.Sequence) (float64, error) {
	bracketSet, err := bracketset.InnerMostBracket(sequence)

	if err != nil {
		return 0.00, err
	}

	if bracketSet.IsEmpty() {
		return calculateSequence(sequence)
	}

	result, err := calculateBracketValue(bracketSet, sequence)

	if err != nil {
		return 0.00, err
	}

	previousSequence := sequence.Previous(bracketSet.Open())
	reminaing := sequence.Remaining(bracketSet.Close())

	return calculateSingleBracketSequence(previousSequence.Join(reminaing, result))
}

func calculateSequence(sequence seq.Sequence) (float64, error) {
	if sequence.IsSingleValue() {
		return strconv.ParseFloat(sequence[0], 10)
	}

	if sequence.IsSimpleSequence() {
		return resolveSimpleSequence(sequence[0], sequence[1], sequence[2])
	}

	for position, segment := range sequence {
		if !isValidOperator(segment) {
			continue
		}

		if isPlusOrMinus(segment) && containsMultiplyOrDivide(sequence) {
			continue
		}

		previous := previous(position)
		next := next(position)

		expression, err := NewSimpleExpression(sequence[previous], sequence[position], sequence[next])

		if err != nil {
			return 0.00, err
		}

		result, err := expression.ResolveToString()

		if err != nil {
			return 0.00, err
		}

		previousSequence := sequence[:previous]
		reminaing := append([]string{result}, sequence[next+1:]...)

		return calculateSequence(append(previousSequence, reminaing...))
	}

	return 00.0, errors.New("Could not resolve sequence")
}

func resolveSimpleSequence(first, operator, second string) (float64, error) {
	expression, err := NewSimpleExpression(first, operator, second)

	if err != nil {
		return 0.00, err
	}

	return expression.Resolve()
}

func calculateBracketValue(bracketSet bracketset.BracketSet, sequence seq.Sequence) (string, error) {
	innerSequence, err := sequence.ReturnInner(bracketSet.Open(), bracketSet.Close())
	if err != nil {
		return "0", err
	}

	result, err := calculateSequence(innerSequence)
	if err != nil {
		return "0", err
	}

	return fmt.Sprintf("%f", result), nil
}

func isMultipleBracketSetSum(sequence seq.Sequence) bool {
	for position, segment := range sequence {
		if sequence.IsFirstOrLast(position) {
			continue
		}

		if isValidOperatorBetweenBrackets(segment, sequence[previous(position)], sequence[next(position)]) {
			return true
		}
	}

	return false
}

func isPlusOrMinus(segment string) bool {
	return strings.Contains("+-", segment)
}

func containsMultiplyOrDivide(sequence seq.Sequence) bool {
	sequenceString := strings.Join(sequence, "")

	return strings.Contains(sequenceString, "/") || strings.Contains(sequenceString, "*")
}

func isValidOperatorBetweenBrackets(segment, before, after string) bool {
	return isValidOperator(segment) && isBracket(before) && isBracket(after)
}

func isBracket(s string) bool {
	return s == "(" || s == ")"
}

func next(position int) int {
	return position + 1
}

func previous(position int) int {
	return position - 1
}
