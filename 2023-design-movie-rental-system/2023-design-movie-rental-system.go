import (
    "fmt"
    "sort"
)

type Entry struct {
    Price int
    Shop  int
}

type RentedEntry struct {
    Price int
    Shop  int
    Movie int
}

type MovieRentingSystem struct {
    // movie -> sorted slice of entries {price, shop}
    unrentedMovies map[int][]Entry
    
    // sorted slice of rented entries {price, shop, movie}
    rentedMovies   []RentedEntry
    
    // price lookup: "shop,movie" -> price
    prices         map[string]int
}

/**
 * Initialize movie rental system
 * TC: O(E log E), SC: O(E)
 */
func Constructor(n int, entries [][]int) MovieRentingSystem {
    system := MovieRentingSystem{
        unrentedMovies: make(map[int][]Entry),
        rentedMovies:   make([]RentedEntry, 0),
        prices:         make(map[string]int),
    }
    
    // Group entries by movie
    for _, entry := range entries {
        shop, movie, price := entry[0], entry[1], entry[2]
        
        system.unrentedMovies[movie] = append(
            system.unrentedMovies[movie],
            Entry{Price: price, Shop: shop},
        )
        
        system.prices[fmt.Sprintf("%d,%d", shop, movie)] = price
    }
    
    // Sort each movie's entries
    for movie := range system.unrentedMovies {
        sort.Slice(system.unrentedMovies[movie], func(i, j int) bool {
            entries := system.unrentedMovies[movie]
            if entries[i].Price != entries[j].Price {
                return entries[i].Price < entries[j].Price
            }
            return entries[i].Shop < entries[j].Shop
        })
    }
    
    return system
}

/**
 * Search cheapest 5 shops for movie
 * TC: O(1), SC: O(1)
 */
func (this *MovieRentingSystem) Search(movie int) []int {
    entries, exists := this.unrentedMovies[movie]
    if !exists {
        return []int{}
    }
    
    result := make([]int, 0, 5)
    for i, entry := range entries {
        if i >= 5 {
            break
        }
        result = append(result, entry.Shop)
    }
    
    return result
}

/**
 * Rent movie from shop
 * TC: O(log n), SC: O(1)
 */
func (this *MovieRentingSystem) Rent(shop int, movie int) {
    priceKey := fmt.Sprintf("%d,%d", shop, movie)
    price := this.prices[priceKey]
    
    // Remove from unrented (binary search + removal)
    entries := this.unrentedMovies[movie]
    target := Entry{Price: price, Shop: shop}
    
    index := sort.Search(len(entries), func(i int) bool {
        if entries[i].Price != target.Price {
            return entries[i].Price >= target.Price
        }
        return entries[i].Shop >= target.Shop
    })
    
    if index < len(entries) && entries[index].Price == price && entries[index].Shop == shop {
        // Remove the entry
        this.unrentedMovies[movie] = append(entries[:index], entries[index+1:]...)
    }
    
    // Add to rented (maintain sorted order)
    rentedEntry := RentedEntry{Price: price, Shop: shop, Movie: movie}
    this.insertSortedRented(rentedEntry)
}

/**
 * Drop rented movie back to shop
 * TC: O(log n), SC: O(1)
 */
func (this *MovieRentingSystem) Drop(shop int, movie int) {
    priceKey := fmt.Sprintf("%d,%d", shop, movie)
    price := this.prices[priceKey]
    
    // Remove from rented
    target := RentedEntry{Price: price, Shop: shop, Movie: movie}
    index := this.findRentedIndex(target)
    if index >= 0 {
        this.rentedMovies = append(this.rentedMovies[:index], this.rentedMovies[index+1:]...)
    }
    
    // Add back to unrented
    entry := Entry{Price: price, Shop: shop}
    this.insertSortedUnrented(movie, entry)
}

/**
 * Report 5 cheapest rented movies
 * TC: O(1), SC: O(1)
 */
func (this *MovieRentingSystem) Report() [][]int {
    result := make([][]int, 0, 5)
    
    for i, entry := range this.rentedMovies {
        if i >= 5 {
            break
        }
        result = append(result, []int{entry.Shop, entry.Movie})
    }
    
    return result
}

/**
 * Helper: Insert rented entry maintaining sorted order
 */
func (this *MovieRentingSystem) insertSortedRented(entry RentedEntry) {
    index := sort.Search(len(this.rentedMovies), func(i int) bool {
        current := this.rentedMovies[i]
        if current.Price != entry.Price {
            return current.Price >= entry.Price
        }
        if current.Shop != entry.Shop {
            return current.Shop >= entry.Shop
        }
        return current.Movie >= entry.Movie
    })
    
    // Insert at the found position
    this.rentedMovies = append(this.rentedMovies, RentedEntry{})
    copy(this.rentedMovies[index+1:], this.rentedMovies[index:])
    this.rentedMovies[index] = entry
}

/**
 * Helper: Insert unrented entry maintaining sorted order
 */
func (this *MovieRentingSystem) insertSortedUnrented(movie int, entry Entry) {
    entries := this.unrentedMovies[movie]
    
    index := sort.Search(len(entries), func(i int) bool {
        if entries[i].Price != entry.Price {
            return entries[i].Price >= entry.Price
        }
        return entries[i].Shop >= entry.Shop
    })
    
    // Insert at the found position
    entries = append(entries, Entry{})
    copy(entries[index+1:], entries[index:])
    entries[index] = entry
    this.unrentedMovies[movie] = entries
}

/**
 * Helper: Find index of rented entry
 */
func (this *MovieRentingSystem) findRentedIndex(target RentedEntry) int {
    for i, entry := range this.rentedMovies {
        if entry.Price == target.Price && entry.Shop == target.Shop && entry.Movie == target.Movie {
            return i
        }
    }
    return -1
}

/**
 * Your MovieRentingSystem object will be instantiated and called as such:
 * obj := Constructor(n, entries);
 * param_1 := obj.Search(movie);
 * obj.Rent(shop,movie);
 * obj.Drop(shop,movie);
 * param_4 := obj.Report();
 */