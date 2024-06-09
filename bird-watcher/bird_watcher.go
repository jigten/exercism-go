package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0

	for _, c := range birdsPerDay {
		total += c
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	startWeek := (week - 1) * 7
	total := 0

	for i := startWeek; i < startWeek+7; i++ {
		total += birdsPerDay[i]
	}

	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	correctedBirdsPerDay := birdsPerDay

	for i := 0; i < len(correctedBirdsPerDay); i += 2 {
		correctedBirdsPerDay[i] += 1
	}

	return correctedBirdsPerDay
}
