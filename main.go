package main

import ()

func main() {

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