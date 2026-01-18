package main

import (
	"errors"
	"fmt"
)

func Plus(a, b int) (int, error) {

	if a > 1000000 || b > 1000000 {
		return 0, errors.New("числа слишком большие")
	}
	return a + b, nil
}

func Minus(a, b int) (int, error) {

	return a - b, nil
}

func Ymnosh(a, b int) (int, error) {

	if a == 0 || b == 0 {
		return 0, errors.New("умножение на ноль невозможно")
	}
	return a * b, nil

}

func Delenie(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("деление на ноль невозможно")
	}
	if a == 0 {
		return 0, errors.New("деление нуля невозможно")
	}
	return a / b, nil
}

func main() {

	fmt.Println("Введите два числа:")

	var a, b int
	fmt.Scan(&a, &b)

	if a > 1000000 || b > 1000000 {
		err := errors.New("числа слишком большие")
		fmt.Println("Ошибка:", err)
	}

	fmt.Println("Поделить - / Умножить * Сложить + Вычесть -")
	var operation string
	fmt.Scan(&operation)

	switch operation {
	case "+":
		result, err := Plus(a, b)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Println("Результат:", result)
	case "-":
		result, err := Minus(a, b)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Println("Результат:", result)
	case "*":
		result, err := Ymnosh(a, b)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Println("Результат:", result)
	case "/":
		result, err := Delenie(a, b)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Println("Результат:", result)
	default:
		fmt.Println("Неправильная операция")
	}
}
