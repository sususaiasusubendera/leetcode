type FoodRatings struct {
	ratingMap  map[string]int       // food -> rating
	cuisineMap map[string]string    // food -> cuisine
	heaps      map[string]*FoodHeap // cuisine -> max heap of foods
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	foodRatings := FoodRatings{
		ratingMap:  make(map[string]int),
		cuisineMap: make(map[string]string),
		heaps:      make(map[string]*FoodHeap),
	}

	for i := 0; i < len(foods); i++ {
		food, cuisine, rating := foods[i], cuisines[i], ratings[i]
		foodRatings.ratingMap[food] = rating
		foodRatings.cuisineMap[food] = cuisine

		newFood := Food{
			food:    food,
			cuisine: cuisine,
			rating:  rating,
		}

		if foodRatings.heaps[cuisine] == nil {
			foodRatings.heaps[cuisine] = &FoodHeap{}
		}

		heap.Push(foodRatings.heaps[cuisine], newFood)
	}

	return foodRatings
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	this.ratingMap[food] = newRating

	cuisine := this.cuisineMap[food]
	newFood := Food{
		food:    food,
		cuisine: cuisine,
		rating:  newRating,
	}

	heap.Push(this.heaps[cuisine], newFood)
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	h := this.heaps[cuisine] // get *FoodHeap of the cuisine
	for { // loop until getting the food with valid rating
        top := (*h)[0] // curr top
        if top.rating == this.ratingMap[top.food] { return top.food }

        heap.Pop(h) // lazy delete the food with invalid rating from the heap
    }
}

// array, hash map, heap
// time: O(log(N))
// space: O(N)

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */

type Food struct {
	food    string
	cuisine string
	rating  int
}

// max heap implementation
type FoodHeap []Food

func (h FoodHeap) Len() int { return len(h) }

func (h FoodHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h FoodHeap) Less(i, j int) bool {
	if h[i].rating == h[j].rating {
		return h[i].food < h[j].food
	}
	return h[i].rating > h[j].rating
}

func (h *FoodHeap) Push(x any) { *h = append(*h, x.(Food)) }

func (h *FoodHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}