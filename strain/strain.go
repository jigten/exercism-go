package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
func Keep[T any](s []T, f func(T) bool) []T {
	res := []T{}

	for _, t := range s {
		if f(t) {
			res = append(res, t)
		}
	}

	return res
}

func Discard[T any](s []T, f func(T) bool) []T {
	res := []T{}

	for _, t := range s {
		if !f(t) {
			res = append(res, t)
		}
	}

	return res
}
