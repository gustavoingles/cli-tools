package main

import (
	"fmt"
	"os"
)

func main() {
	calculator := calculator{}
	fmt.Println("If you want to sum two numbers, type \"sum\", if you want to subtract two numbers, type \"subtract\", if you want to multiply two numbers, type \"multiply\", if you want to divide two numbers, type \"divide\"")
	if os.Args[1] == "sum" {
		for i, _ := range os.Args[2:] {
			number := float64(i)
			calculator.firstNumber = number
		}
	}
}

type calculator struct {
	firstNumber float64
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