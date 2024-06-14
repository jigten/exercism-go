package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	steps := 0

	if n <= 0 {
		return steps, errors.New("zero or negative numbers not allowed")
	}

	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}

		steps++
	}

	return steps, nil
}
