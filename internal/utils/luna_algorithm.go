package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func LunaCheckInt(data int) (bool, error) {
	digits := convertIntToSliceOfIntegers(data)
	return lunaCheck(digits), nil
}

func LunaCheckString(data string) (bool, error) {
	match, err := regexp.MatchString(`^\d+$`, data)
	switch {
	case data == "":
		return false, errors.New("empty data provided")
	case err != nil:
		return false, err
	case !match:
		return false, fmt.Errorf("строка '%s' содержит недопустимые символы", data)
	}

	digits := convertStringToSliceOfIntegers(data)

	return lunaCheck(digits), nil
}

func lunaCheck(input []int) bool {
	if input == nil {
		return false
	}

	var sum int
	parity := len(input) % 2
	for i := 0; i < len(input); i++ {
		digit := input[i]
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}

func convertStringToSliceOfIntegers(data string) []int {
	var digits []int
	for _, r := range data {
		d, _ := strconv.Atoi(string(r))
		digits = append(digits, d)
	}

	return digits
}

func convertIntToSliceOfIntegers(data int) []int {
	buf := make([]byte, 0, 20)
	buf = strconv.AppendInt(buf, int64(data), 10) // 0 allocations

	digits := make([]int, 0, len(buf))
	for _, b := range buf {
		digits = append(digits, int(b-'0'))
	}

	return digits
}
