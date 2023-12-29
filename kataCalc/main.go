package main

import (
	"fmt"
	"strconv"
	"strings"
)

const operators = "+-*/"

func main() {
	var exp string

	_, err := fmt.Scan(&exp)
	if err != nil {
		return
	}

	first, second, operator := parseExpression(exp)
	fmt.Println(parseNumbers(first, second, operator))
}

func parseOperator(exp string) (operator string) {
	for _, v := range exp {
		if strings.ContainsRune(operators, v) {
			operator += string(v)
		}
	}

	if len(operator) != 1 {
		if len(operator) == 0 {
			panic("Вывод ошибки, так как строка не является математической операцией.")
		}

		panic("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	return
}

func parseExpression(exp string) (first, second, operator string) {
	operator = parseOperator(exp)
	operands := strings.Split(exp, operator)

	first = operands[0]
	second = operands[1]

	return
}

func parseNumbers(first, second, operator string) string {
	firstNum, err1 := strconv.Atoi(first)
	secondNum, err2 := strconv.Atoi(second)

	if err1 == nil && err2 == nil {
		return strconv.Itoa(calculate(firstNum, secondNum, operator))
	}

	if err1 != nil && err2 != nil {
		return parseRome(first, second, operator)
	}

	panic("Вывод ошибки, так как используются одновременно разные системы счисления.")
}

func calculate(first, second int, operator string) (result int) {
	if first < 11 && second < 11 {
		switch operator {
		case "+":
			result = first + second
		case "-":
			result = first - second
		case "*":
			result = first * second
		case "/":
			result = first / second
		}
		return
	} else {
		panic("Вывод ошибки, так как используются числа больше 10.")
	}
}

func parseRome(first, second, operator string) (result string) {
	romeNums := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	firstRome, ok1 := romeNums[first]
	secondRome, ok2 := romeNums[second]

	if ok1 && ok2 {
		resultNum := calculate(firstRome, secondRome, operator)

		if resultNum < 0 {
			panic("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
		}

		symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

		for i := 0; i < len(symbols); i++ {
			for resultNum >= values[i] {
				result += symbols[i]
				resultNum -= values[i]
			}
		}

		return
	}

	panic("Вывод ошибки, так как введенная информация не является арабским или римским числом.")
}
