package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operator string

var romanToArabicMap = map[string]int{
	"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10,
}

var arabicToRomanMap = []struct {
	Value  int
	Symbol string
}{
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func romanToArabic(roman string) (int, error) {
	arabic := 0
	for i := 0; i < len(roman); {
		if i+1 < len(roman) && romanToArabicMap[roman[i:i+2]] > 0 {
			arabic += romanToArabicMap[roman[i:i+2]]
			i += 2
		} else if romanToArabicMap[string(roman[i])] > 0 {
			arabic += romanToArabicMap[string(roman[i])]
			i++
		} else {
			return 0, errors.New("некорректное римское число")
		}
	}
	return arabic, nil
}

func arabicToRoman(arabic int) string {
	var roman strings.Builder
	for _, item := range arabicToRomanMap {
		for arabic >= item.Value {
			roman.WriteString(item.Symbol)
			arabic -= item.Value
		}
	}
	return roman.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.ReplaceAll(strings.TrimSpace(input), " ", "")

	var operator operator
	switch {
	case strings.ContainsAny(input, "*"):

		operator = "*"
	case strings.ContainsAny(input, "/"):

		operator = "/"
	case strings.ContainsAny(input, "+"):

		operator = "+"
	case strings.ContainsAny(input[1:], "-"):

		operator = "-"
	default:
		panic("Ошибка: неизвестный оператор")
	}

	var parts []string
	parts = strings.Split(input, string(operator))
	if len(parts) != 2 {
		panic("Ошибка: неправильный формат выражения")
	}

	isRoman := false
	fNum, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
	sNum, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err1 != nil && err2 != nil {
		fNum, err1 = romanToArabic(strings.TrimSpace(parts[0]))
		sNum, err2 = romanToArabic(strings.TrimSpace(parts[1]))
		if err1 != nil || err2 != nil {
			panic("Ошибка преобразования чисел")
		}
		isRoman = true
	} else if err1 != nil || err2 != nil {
		panic("Ошибка: нельзя смешивать арабские и римские цифры")
	}

	if fNum < 1 || fNum > 10 || sNum < 1 || sNum > 10 {
		panic("Ошибка: числа должны быть в диапазоне от 1 до 10 включительно")
	}

	var result int
	switch operator {
	case "+":
		result = fNum + sNum
	case "-":
		result = fNum - sNum
	case "*":
		result = fNum * sNum
	case "/":
		result = fNum / sNum
	}

	if isRoman {
		if result < 1 {
			panic("Ошибка: результат вычисления римских чисел не может быть меньше единицы")
		}
		fmt.Println("Результат:", arabicToRoman(result))
	} else {
		fmt.Println("Результат:", result)
	}
}
