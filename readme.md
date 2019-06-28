# Basic Calculator in Go

This a basic calculator module with cli applications that can preform basic arithmetic.

The calculator adheres to BoDMAS rules. It can only perform Multiplication, Division, Addition and Subtraction.

## Intentions

This project was developed to help practice and expand my skills in the Go language, as well as an opportunity to learn how the GO ecosystem and workflow works.


## Motivation

This is one of the first pieces of Go-lang code I have written outside of basic tutorials and welcome in criticism and feedback. The goal of this project was to practice TDD in GO.
 

## How it works

The calculator works by taking in a string sum returning a float with the answer or an error.

A sum can consist of Multiplication, Division, Addition and Subtraction and brackets. like below

`"2 + 3 * 4 / 4 - 6"`

The calculator can be used as a GO module inside other projects or as compiled CLI application.


## Install 

From the project root run:

`go install`

The calculator application can now be run from `goworkspace/bin/basiccalculator`


### Using Cli application

Once installed the calculator can be used in the command line by the following.

`./basiccalculator "2+2"`

This will return `4.0`

The sum provided has to be in quotes and given as a string value. Spaces can be left out or included

The following are all treated as same "`2 * 4`", "`2* 4`" and "`2*4`"

## Using as Module

The module can be used with the standard `go get` functionality. To use the calculator simply pull the `calculatorservice` into your project with the following.

```import "github.com/kasch22/basiccalculator/calculatorservice"```

The Calculator Service has two exported functions. The primary one is the `CalculateString` function. This takes in the sum as a string and converts it to a `calculatorservice/sequence.Sequence` type before calculating the value of the sum.

The `CreateSequence` function is available from the `github.com/kasch22/basiccalculator/calculatorservice/sequence` module. This will convert a string sum into a Sequence type should you wish to use this type before it being calculated.
