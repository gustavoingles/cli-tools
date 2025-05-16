package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"log"
)

type calculator struct {
	firstNumber  float64
	secondNumber float64
}

func initializeCalculator() calculator {
	fmt.Println("If you want to sum two numbers, type \"sum\", if you want to subtract two numbers, \ntype \"subtract\", if you want to multiply two numbers, type \"multiply\", if you want to divide two numbers, type \"divide\"")
	calculator := calculator{}

	return calculator
}

func (c *calculator) calculatorWorkflow() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	chosenOperation := scanner.Text()
	if chosenOperation == "sum" {
		err := c.collectNumbers()
		if err!= nil {
			log.Printf("Failed to collect numbers: %v", err)
		}

		totalSum := c.sum()

		fmt.Printf("The sum between %.2f and %.2f is equals to %.2f\n", c.firstNumber, c.secondNumber, totalSum)
	} else if chosenOperation == "subtract" {
		err := c.collectNumbers()
			if err!= nil {
				log.Printf("Failed to collect numbers: %v", err)
			}

		totalSubtraction := c.subtract()

		fmt.Printf("The subtraction between %.2f and %.2f is equals to %.2f\n", c.firstNumber, c.secondNumber, totalSubtraction)
	} else if chosenOperation == "multiply" {
		err := c.collectNumbers()
			if err!= nil {
				log.Printf("Failed to collect numbers: %v", err)
			}

		totalMultiplication := c.multiply()

		fmt.Printf("The subtraction between %.2f and %.2f is equals to %.2f\n", c.firstNumber, c.secondNumber, totalMultiplication)
	} else if chosenOperation == "divide" {
		err := c.collectNumbers()
		if err!= nil {
			log.Printf("Failed to collect numbers: %v", err)
		}

		totalDivision, err := c.divide()
		if err != nil {
			log.Fatalln("You can't divide any number by zero, please restart the program and choose another combination of numbers.")
		}

		fmt.Printf("The subtraction between %.2f and %.2f is equals to %.2f\n", c.firstNumber, c.secondNumber, totalDivision)
	}
}

func (c *calculator) collectNumbers() error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please choose the first number: ")
	scanner.Scan()
	firstNumber := scanner.Text()
	convFirstNumber, err := strconv.ParseFloat(firstNumber, 64)
	if err != nil {
		return fmt.Errorf("it's happened an error while parsing the first number, please verify it: %w", err)
	}
	c.firstNumber = convFirstNumber

	fmt.Print("Please choose the second number: ")
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
