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
		scanner.Scan()
		firstNumber := scanner.Text()
		convFirstNumber, err := strconv.ParseFloat(firstNumber, 64)
		if err != nil {
			log.Printf("It's happened an error while parsing the chosen number, please verify it: %v", err)
		}
		calculator.firstNumber = convFirstNumber

		scanner.Scan()
		secondNumber := scanner.Text()
		convSecondNumber, err := strconv.ParseFloat(secondNumber, 64)
		if err != nil {
			log.Printf("It's happened an error while parsing the chosen number, please verify it: %v", err)
		}
		calculator.secondNumber = convSecondNumber

		totalSum := calculator.firstNumber + calculator.secondNumber

		fmt.Printf("The sum between %f and %f is equals to %f\n", calculator.firstNumber, calculator.secondNumber, totalSum)
	} else if chosenOperation == "subtract" {
		scanner.Scan()
		firstNumber := scanner.Text()
		convFirstNumber, err := strconv.ParseFloat(firstNumber, 64)
		if err != nil {
			log.Printf("It's happened an error while parsing the chosen number, please verify it: %v", err)
		}
		calculator.firstNumber = convFirstNumber

		scanner.Scan()
		secondNumber := scanner.Text()
		convSecondNumber, err := strconv.ParseFloat(secondNumber, 64)
		if err != nil {
			log.Printf("It's happened an error while parsing the chosen number, please verify it: %v", err)
		}
		calculator.secondNumber = convSecondNumber

		totalSubtraction := calculator.firstNumber - calculator.secondNumber

		fmt.Printf("The subtraction between %f and %f is equals to %f\n", calculator.firstNumber, calculator.secondNumber, totalSubtraction)
	}
}

type calculator struct {
	firstNumber  float64
	secondNumber float64
}

type operations interface {
	sum(x, y float64) float64
	subtract(x, y float64) float64
	multiply(x, y float64) float64
	divide(x, y float64) float64
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

func (c calculator) divide() float64 {
	return c.firstNumber / c.secondNumber
}
