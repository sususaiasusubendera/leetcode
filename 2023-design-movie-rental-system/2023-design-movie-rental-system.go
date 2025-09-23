type Entry struct {
	shop  int
	movie int
	price int
}

type MovieRentingSystem struct {
	availMovieMap map[int]*AvailableHeap // movie -> AvailableHeap
	priceMap      map[[2]int]int         // shop, movie -> price
	rented        *RentedHeap
	rentedMap     map[[2]int]struct{} // shop, movie -> struct{} (lazy delete)
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	mrs := MovieRentingSystem{
		availMovieMap: make(map[int]*AvailableHeap),
		priceMap:      make(map[[2]int]int),
		rented:        &RentedHeap{},
		rentedMap:     make(map[[2]int]struct{}),
	}

	for _, entry := range entries {
		shop, movie, price := entry[0], entry[1], entry[2]
		newEntry := Entry{
			shop:  shop,
			movie: movie,
			price: price,
		}

		mrs.priceMap[[2]int{shop, movie}] = price

		if mrs.availMovieMap[movie] == nil {
			mrs.availMovieMap[movie] = &AvailableHeap{}
		}

		heap.Push(mrs.availMovieMap[movie], newEntry)
	}

	return mrs
}

func (this *MovieRentingSystem) Search(movie int) []int {
	h := this.availMovieMap[movie]
	res := []int{}
	if h == nil || h.Len() == 0 {
		return res
	}

	temp := []Entry{}
	seenShop := make(map[int]struct{}) // prevent duplicate (shop, movie)
	for len(res) < 5 && h.Len() > 0 {
		top := heap.Pop(h).(Entry)
		pair := [2]int{top.shop, top.movie}

        // lazy delete if it is rented (unavailable)
		if _, rented := this.rentedMap[pair]; rented { continue }
       
        // lazy delete if it appears in the results
		if _, seen := seenShop[top.shop]; seen { continue }

		res = append(res, top.shop)
		seenShop[top.shop] = struct{}{}
		temp = append(temp, top)
	}

	// push back all valid entries
	for _, e := range temp {
		heap.Push(h, e)
	}

	return res
}

func (this *MovieRentingSystem) Rent(shop int, movie int) {
	price := this.priceMap[[2]int{shop, movie}]
	entry := Entry{shop: shop, movie: movie, price: price}

	this.rentedMap[[2]int{shop, movie}] = struct{}{}
	heap.Push(this.rented, entry)
}

func (this *MovieRentingSystem) Drop(shop int, movie int) {
	price := this.priceMap[[2]int{shop, movie}]
	entry := Entry{shop: shop, movie: movie, price: price}

	delete(this.rentedMap, [2]int{shop, movie})
	heap.Push(this.availMovieMap[movie], entry)
}

func (this *MovieRentingSystem) Report() [][]int {
	h := this.rented
	res := [][]int{}
	if h == nil || h.Len() == 0 {
		return res
	}

	temp := []Entry{}
	seen := make(map[[2]int]struct{}) // prevent duplicate (shop, movie)
	for len(res) < 5 && h.Len() > 0 {
		top := heap.Pop(h).(Entry)
		pair := [2]int{top.shop, top.movie}

        // lazy delete if it isn't rented
		if _, rented := this.rentedMap[pair]; !rented { continue }

        // lazy delete if it appears in the results
		if _, ok := seen[pair]; ok { continue }

		r := []int{top.shop, top.movie}
		res = append(res, r)
		seen[pair] = struct{}{}
		temp = append(temp, top)
	}

	// push back all valid entries
	for _, e := range temp {
		heap.Push(h, e)
	}

	return res
}

// array, design, hash map, heap
// time:
// - Constructor: O(nlog(n))
// - Search: amortized O(log(n))
// - Rent: O(log(n))
// - Drop: O(lof(n))
// - Report: amortized O(log(n))
// space:
// - Constructor: O(n)
// - Search: amortized O(n)
// - Rent: O(1)
// - Report: amortized O(n)

/**
 * Your MovieRentingSystem object will be instantiated and called as such:
 * obj := Constructor(n, entries);
 * param_1 := obj.Search(movie);
 * obj.Rent(shop,movie);
 * obj.Drop(shop,movie);
 * param_4 := obj.Report();
 */

// AvailableHeap implementation
type AvailableHeap []Entry

func (h *AvailableHeap) Len() int { return len(*h) }

func (h *AvailableHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *AvailableHeap) Less(i, j int) bool {
	if (*h)[i].price == (*h)[j].price {
		return (*h)[i].shop < (*h)[j].shop
	}

	return (*h)[i].price < (*h)[j].price
}

func (h *AvailableHeap) Push(x any) { *h = append(*h, x.(Entry)) }

func (h *AvailableHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

// RentedHeap implementation
type RentedHeap []Entry

func (h *RentedHeap) Len() int { return len(*h) }

func (h *RentedHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *RentedHeap) Less(i, j int) bool {
	if (*h)[i].price == (*h)[j].price && (*h)[i].shop == (*h)[j].shop {
		return (*h)[i].movie < (*h)[j].movie
	}

	if (*h)[i].price == (*h)[j].price {
		return (*h)[i].shop < (*h)[j].shop
	}

	return (*h)[i].price < (*h)[j].price
}

func (h *RentedHeap) Push(x any) { *h = append(*h, x.(Entry)) }

func (h *RentedHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}