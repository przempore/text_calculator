package main

import (
	"strconv"
	"strings"
)

func main() {

}

func Calculate(calc string) int {
	calc = strings.Replace(calc, " ", "", -1)
	var digits []int
	var signs []string
	start_number := -1
	stop_number := -1
	is_digit_started := true
	for i, char := range calc {
		switch string(char) {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			if is_digit_started {
				start_number = i;
			}
			is_digit_started = false
		case "+":
			stop_number = i
			digit, _ := get_number(start_number, stop_number, calc)
			digits = append(digits, digit)
			signs = append(signs, string(char))
			is_digit_started = true
		}
	}

	digit, _ := get_number(start_number, len(calc), calc)
	digits = append(digits, digit)

	return do_calculations(digits[0], digits[1], signs[0])
}
func get_number(start int, stop int, calc string) (int, error) {
	number := calc[start:stop]
	return strconv.Atoi(number)
}

func do_calculations(x, y int, sign string) int {
	switch string(sign) {
	case "+":
		return addition(x, y)
	}
	return -1
}

func addition(x int, y int) int {
	return x + y
}
