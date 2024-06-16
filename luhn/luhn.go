package luhn

import (
	"strconv"
)

func Valid(id string) bool {
	newId := []rune(id)
	digRev := []int{}

	for i := len(newId) - 1; i >= 0; i-- {
		if newId[i] == ' ' {
			continue
		}

		n, err := strconv.Atoi(string(newId[i]))

		if err != nil {
			return false
		}

		digRev = append(digRev, n)
	}

	if len(digRev) < 2 {
		return false
	}

	sum := 0

	for i, n := range digRev {
		if i%2 != 0 {
			n *= 2

			if n > 9 {
				n -= 9
			}
		}

		sum += n
	}

	return sum%10 == 0
}
