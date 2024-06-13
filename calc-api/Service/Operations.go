package Service

import "fmt"

type Operations interface {
	Add(operand1, operand2 float64) float64
	Subtract(operand1, operand2 float64) float64
	Divide(operand1, operand2 float64) (float64, error)
	Multiply(operand1, operand2 float64) float64
}

type operation struct{}

func (operation) Add(operand1, operand2 float64) float64 {
	return operand1 + operand2
}

func (operation) Subtract(operand1, operand2 float64) float64 {
	return operand1 - operand2
}

func (operation) Multiply(operand1, operand2 float64) float64 {
	return operand1 * operand2
}

func (operation) Divide(operand1, operand2 float64) (float64, error) {
	if operand2 == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return operand1 / operand2, nil
}

func CreateOperationInstance() Operations {
	return &operation{}
}
