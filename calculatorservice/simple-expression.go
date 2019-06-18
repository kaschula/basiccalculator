package calculatorservice

// ^^ change the package name sequence service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func NewSimpleExpression(first, operator, second string) (SimpleExpression, error) {
	firstDigit, err := strconv.ParseFloat(first, 10)
	if err != nil {
		return SimpleExpression{}, err
	}

	secondDigit, err := strconv.ParseFloat(second, 10)
	if err != nil {
		return SimpleExpression{}, err
	}

	resolvedOperator, err := parseOperator(operator)
	if err != nil {
		return SimpleExpression{}, err
	}

	return SimpleExpression{firstDigit, resolvedOperator, secondDigit}, nil
}

type SimpleExpression struct {
	firstDigit  float64
	operator    string
	secondDigit float64
}

func (se SimpleExpression) Resolve() (float64, error) {
	//TODO: change strings to consts
	switch se.operator {
	case "*":
		return se.firstDigit * se.secondDigit, nil
	case "/":
		return se.firstDigit / se.secondDigit, nil
	case "+":
		return se.firstDigit + se.secondDigit, nil
	case "-":
		return se.firstDigit - se.secondDigit, nil
	default:
		return 0, errors.New(fmt.Sprintf("Error Occurred, can/'t calculate %v", se))
	}
}

func (se SimpleExpression) ResolveToString() (string, error) {
	result, err := se.Resolve()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%f", result), nil
}

func parseOperator(operator string) (string, error) {
	if !isValidOperator(operator) {
		return "", errors.New(fmt.Sprintf("Error: Unknown Operator %v", operator))
	}

	return operator, nil
}

// This should be pulled out into own file
func isValidOperator(operator string) bool {
	validOperators := "/*+-"

	return strings.Contains(validOperators, operator)
}
