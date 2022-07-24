package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int
	fmt.Printf("enter a number: ")
	fmt.Scanf("%d", &number)
	fmt.Println(factorial(number))
}

func factorial(number int) (int, error) {
	//base case
	if number == 1 || number == 0 {
		return 1, nil
	} else if number < 0 {
		return number, errors.New("number cannot be negative")
	}
	n, err := factorial(number - 1)
	return number * n, err
}
