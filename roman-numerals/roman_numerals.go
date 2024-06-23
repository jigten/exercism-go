package romannumerals

import (
	"errors"
	"strings"
)

var romanToDec = []struct {
	decimal int
	roman   string
}{
	{1, "I"}, {4, "IV"}, {5, "V"}, {9, "IX"},
	{10, "X"}, {40, "XL"}, {50, "L"}, {90, "XC"},
	{100, "C"}, {400, "CD"}, {500, "D"}, {900, "CM"},
	{1000, "M"},
}

func ToRomanNumeral(dec int) (string, error) {
	if dec <= 0 || 4000 <= dec {
		return "", errors.New("invalid range")
	}

	roman := ""
	n := 0

	for i := len(romanToDec) - 1; 0 <= i; i-- {
		n, dec = divMod(dec, romanToDec[i].decimal)
		roman += strings.Repeat(romanToDec[i].roman, n)
	}
	return roman, nil
}

func divMod(numerator, divisor int) (int, int) {
	return numerator / divisor, numerator % divisor
}
