package calculator

import "errors"

func Calculate(operation string, FirstOperand, SecondOperand int) (result int, err error) {

	if operation == "add" {

		result = FirstOperand + SecondOperand

		return result, nil

	} else if operation == "multiply" {

		result = FirstOperand * SecondOperand

		return result, nil

	} else if operation == "substract" {

		result = FirstOperand - SecondOperand

		return result, nil

	} else if operation == "divide" {

		if SecondOperand == 0 {

			return 0, errors.New("Cannot divide by zero... ¿!Are you trying to destroy the world!?")

		} else {

			result := FirstOperand / SecondOperand

			return result, nil
		}

	} else {

		return 0, errors.New("Not a valid operation")

	}

}
