package lasagna

func PreparationTime(input []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return len(input) * prepTime
}

func Quantities(input []string) (noodles int, sauce float64) {
	for _, s := range input {
		if s == "noodles" {
			noodles += 50
		}
		if s == "sauce" {
			sauce += 0.2
		}
	}

	return
}

func AddSecretIngredient(friend []string, mine []string) {
	mine[len(mine)-1] = friend[len(friend)-1]
}

func ScaleRecipe(input []float64, portions int) []float64 {
	result := make([]float64, len(input))
	for i := range input {
		result[i] = input[i] / 2 * float64(portions)
	}

	return result
}
