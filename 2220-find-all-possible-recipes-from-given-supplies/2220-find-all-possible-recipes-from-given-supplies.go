func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	book := make(map[string][]string)
	for i, r := range recipes {
		book[r] = ingredients[i]
	}
	fridge := make(map[string]struct{})
	for _, s := range supplies {
		fridge[s] = struct{}{}
	}
	cooked := []string{}
	batch := recipes
	prev := -1
	for len(cooked) != prev {
		prev = len(cooked)
		next := []string{}
	NextRecipe:
		for _, r := range batch {
			for _, ing := range book[r] {
				if _, ok := fridge[ing]; !ok {
					next = append(next, r)
					continue NextRecipe
				}
			}
			fridge[r] = struct{}{}
			cooked = append(cooked, r)
		}
		batch = next
	}
	return cooked
}

// NOTICE ME SENPAI!!!