// Package lasagna holds tools to help with preparing and cooking a lasagna
package lasagna

const (
	sauce   = "sauce"
	noodles = "noodles"
)

var layerQuantities = map[string]float64{
	sauce:   0.2,
	noodles: 50.0,
}

// PreparationTime estimates the total time in mins to prep based on the amount of layers
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = 2
	}
	return len(layers) * avgPrepTime
}

// Quantities calculates the total amount of certain ingredients needed for the lasagna
func Quantities(layers []string) (int, float64) {
	total := map[string]float64{
		sauce:   0.0,
		noodles: 0.0,
	}
	for _, layer := range layers {
		if quantity, found := layerQuantities[layer]; found {
			total[layer] += quantity
		}
	}
	return int(total[noodles]), total[sauce]
}

// AddSecretIngredient returns a list containing myList and the last ingredient from friendsList
func AddSecretIngredient(friendsList []string, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

// ScaleRecipe returns the amounts of ingredients needed for more portions
func ScaleRecipe(twoPortionAmounts []float64, portions int) []float64 {
	scaled := make([]float64, len(twoPortionAmounts))
	for i, ingredient := range twoPortionAmounts {
		scaled[i] = ingredient / 2 * float64(portions)
	}
	return scaled
}
