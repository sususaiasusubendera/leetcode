func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    count := 0
    for i := 0; i < len(arr1); i++ {
        valid := true
        for j := 0; j < len(arr2); j++ {
            if abs(arr1[i]-arr2[j]) <= d {
                valid = false
                break
            }
        }
        if valid {
            count++
        }
    }
    return count
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

// array
// time: O(nm)
// space: O(1)