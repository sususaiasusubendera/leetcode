func sortByBits(arr []int) []int {
    sort.Slice(arr, func(i, j int) bool {
        if countSetBits(arr[i]) == countSetBits(arr[j]) {
            return arr[i] < arr[j]
        }

        return countSetBits(arr[i]) < countSetBits(arr[j])
    })
    return arr
}

// array, bit manipulation, sorting
// time: O(nlog(n) * log(M))
// space: O(1)

// brian kernighan’s algorithm
func countSetBits(n int) int {
    count := 0
    for n > 0 {
        n &= (n - 1)
        count++
    }
    return count
} 