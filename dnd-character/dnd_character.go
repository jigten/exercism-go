package dndcharacter

import (
	"math"
	"math/rand"
	"slices"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	score := 0
	rolls := []int{}

	for i := 0; i < 4; i++ {
		currentRoll := rand.Intn(7)
		if currentRoll == 0 {
			currentRoll += 1
		}
		rolls = append(rolls, currentRoll)
	}

	minRoll := slices.Min[[]int](rolls)

	for _, r := range rolls {
		score += r
	}

	return score - minRoll
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	str := Ability()
	dex := Ability()
	con := Ability()
	intel := Ability()
	wis := Ability()
	cha := Ability()
	hit := 10 + Modifier(con)

	return Character{str, dex, con, intel, wis, cha, hit}
}
