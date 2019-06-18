# Basic Calculator in Go

This is a basic calculator module with cli to perform basic arithmetic.

The calculator adheres to BoDMAS rules. It can only perform Multiplication, Division, Addition and Subtraction.

## Intentions

This project was developed to help practice and expand my skills in the Go Lang as well as an opportunity to learn how the Go ecosystem and workflow works.

## How it works

The calculator works by taking in a string sum returning a float with the answer or an error.

A sum can consist of Multiplication, Division, Addition and Subtraction and brackets. like below

e.g copy examples from the tests to give examples

The calculator can be used as go module inside othe projects or as compiled CLI application.

### Using as Module

The module can be used with the standard `go get` functionality. To use the calculator simple pull the `calculatorservice` into your project with the following.

```import "github.com/kasch22/basiccalculator/calculatorservice"```

The Calculator Service has two exported functions. The primary one is the `CalculateString` function. This takes in the sum as a string and converts it to a `calculatorservice/sequence.Sequence` type before calculating the value of the sum.

e.g .....

The `CreateSequence` function is available from the `github.com/kasch22/basiccalculator/calculatorservice/sequence` module. This will convert a string sum into a Sequence type should you wish to use this type before it being calculated

e.g using sequence factory

and then calculate by

e.g calculating a sequence


### Using Cli application

The CLI interface is very small and combines some of the examples above. Simple install the calculator

e.g. go install to create the binary 

then in the command line type

e.g using the command line to calculate sums


## Other

This is one of the first pieces of go code I have written outside of basic tutorials and welcome in criticism and feedback.
 