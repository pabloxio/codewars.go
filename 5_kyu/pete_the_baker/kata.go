// https://www.codewars.com/kata/525c65e51bf619685c000059/train/go
package kata

func Min(s []int) int {
	lower := s[0]

	for _, num := range s {
		if num < lower {
			lower = num
		}
	}

	return lower
}

func Cakes(recipe, available map[string]int) int {
	possibleCakes := make([]int, 0)

	for ingredient, requiredAmount := range recipe {
		availableIngredient, ok := available[ingredient]
		if !ok {
			return 0
		}

		possibleCakes = append(possibleCakes, int(availableIngredient/requiredAmount))
	}

	return Min(possibleCakes)
}
