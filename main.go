package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	calculator := calculator{}
	fmt.Println("If you want to sum two numbers, type \"sum\", if you want to subtract two numbers, type \"subtract\", if you want to multiply two numbers, type \"multiply\", if you want to divide two numbers, type \"divide\"")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	chosenOperation := scanner.Text()
	if chosenOperation == "sum" {
		err := calculator.collectNumbers()
		if err!= nil {
			log.Printf("Failed to collect numbers: %v", err)
		}

		totalSum := calculator.sum()

		fmt.Printf("The sum between %f and %f is equals to %f\n", calculator.firstNumber, calculator.secondNumber, totalSum)
	} else if chosenOperation == "subtract" {
		err := calculator.collectNumbers()
			if err!= nil {
				log.Printf("Failed to collect numbers: %v", err)
			}

		totalSubtraction := calculator.subtract()

		fmt.Printf("The subtraction between %f and %f is equals to %f\n", calculator.firstNumber, calculator.secondNumber, totalSubtraction)
	} else if chosenOperation == "multiply" {
		err := calculator.collectNumbers()
			if err!= nil {
				log.Printf("Failed to collect numbers: %v", err)
			}

		totalMultiplication := calculator.multiply()

		fmt.Printf("The subtraction between %f and %f is equals to %f\n", calculator.firstNumber, calculator.secondNumber, totalMultiplication)
	} else if chosenOperation == "divide" {
		err := calculator.collectNumbers()
		if err!= nil {
			log.Printf("Failed to collect numbers: %v", err)
		}

		totalDivision, err := calculator.divide()
		if err != nil {
			fmt.Println("You can't divide any number by zero, please restart the program and choose another combination of numbers, please")
		}

		fmt.Printf("The subtraction between %f and %f is equals to %f\n", calculator.firstNumber, calculator.secondNumber, totalDivision)
	}
}

type calculator struct {
	firstNumber  float64
	secondNumber float64
}

func (c *calculator) collectNumbers() error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please choose the first number")
	scanner.Scan()
	firstNumber := scanner.Text()
	convFirstNumber, err := strconv.ParseFloat(firstNumber, 64)
	if err != nil {
		return fmt.Errorf("it's happened an error while parsing the first number, please verify it: %w", err)
	}
	c.firstNumber = convFirstNumber

	fmt.Println("Please choose the second number")
	scanner.Scan()
	secondNumber := scanner.Text()
	convSecondNumber, err := strconv.ParseFloat(secondNumber, 64)
	if err != nil {
		return fmt.Errorf("it's happened an error while parsing the second number, please verify it: %w", err)
	}
	c.secondNumber = convSecondNumber

	return nil
}

func (c calculator) sum() float64 {
	return c.firstNumber + c.secondNumber
}

func (c calculator) subtract() float64 {
	return c.firstNumber - c.secondNumber
}

func (c calculator) multiply() float64 {
	return c.firstNumber * c.secondNumber
}

func (c calculator) divide() (float64, error) {
	if c.secondNumber == 0 {
		return 0, fmt.Errorf("it's impossible to divide any number by zero, please choose other number")
	}
	return c.firstNumber / c.secondNumber, nil
}
