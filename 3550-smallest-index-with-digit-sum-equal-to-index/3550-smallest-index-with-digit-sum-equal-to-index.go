func smallestIndex(nums []int) int {
    for i, num := range nums {
        if i == sumDigits(num) {
            return i
        }
    }
    return -1
}

func sumDigits(n int) int {
    sum := 0
    for n > 0 {
        sum += n % 10
        n /= 10
    }
    return sum
}

// array, math
// time: O(nd)
// space: O(1)