package main

import (
	"fmt"
	"os"

	"github.com/kasch22/basiccalculator/calculatorservice"
)

func main() {
	sum := os.Args[1]
	value, err := calculatorservice.CalculateString(sum)

	handleError(err)

	fmt.Printf("%#v = %6f", sum, value)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
