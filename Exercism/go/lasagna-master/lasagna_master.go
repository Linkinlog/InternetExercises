package lasagna

// PreparationTime takes a string of layers, and the average time it will take per layer
// and returns an estimate on the total time to do all layers.
func PreparationTime(layers []string, avgTime int) (totalTimeEstimate int) {
	if avgTime == 0 {
		avgTime = 2
	}
	return len(layers) * avgTime
}

// Quantities takes a string of layers and returns the amount of noodles,
// and sauce that we will be needing.
func Quantities(layers []string) (noodles int, sauce float64) {
	noodMultiplier := 50
	sauceMultiplier := 0.2
	for _, layer := range layers {
		if layer == "noodles" {
			noodles = noodles + noodMultiplier
		} else if layer == "sauce" {
			sauce = sauce + sauceMultiplier
		}
	}
	return noodles, sauce
}

// AddSecretIngredient gets the secret ingredient from the last item of stringSlice1,
// and sets it to the last item of stringSlice2 which should always be "?".
func AddSecretIngredient(stringSlice1, stringSlice2 []string) {
	stringSlice2[len(stringSlice2)-1] = stringSlice1[len(stringSlice1)-1]
}

// ScaleRecipe will take a slice of quantities needed for 2 portions and the amount of portions.
// Returns a slice of the scaled quantities that will make the portions requested.
func ScaleRecipe(quantities []float64, portions int) (scaledQuantities []float64) {
	for _, el := range quantities {
		scaledQuantities = append(scaledQuantities, (el/2)*float64(portions))
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
