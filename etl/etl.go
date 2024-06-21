package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	transformed := map[string]int{}

	for k, v := range in {
		for _, s := range v {
			transformed[strings.ToLower(s)] = k
		}
	}

	return transformed
}
