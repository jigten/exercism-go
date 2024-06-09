package lasagna

func PreparationTime(layers []string, preparationTime int) int {
	if preparationTime <= 0 {
		return len(layers) * 2
	}

	return len(layers) * preparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		}

		if layer == "sauce" {
			sauce += 0.2
		}
	}

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numOfPortions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))

	for i := 0; i < len(scaledQuantities); i++ {
		scaledQuantities[i] = quantities[i] * float64(numOfPortions) / 2
	}

	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
