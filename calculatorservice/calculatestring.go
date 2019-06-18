package calculatorservice

import (
	"errors"

	seq "github.com/kasch22/basiccalculator/calculatorservice/sequence"
)

func CalculateString(sum string) (float64, error) {
	if len(sum) == 0 {
		return 0.00, errors.New("No Sum Given")
	}

	sequence, err := seq.CreateSequence(sum)
	if err != nil {
		return 0.00, err
	}

	return Calculate(sequence)
}
