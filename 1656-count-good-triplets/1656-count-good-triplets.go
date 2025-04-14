func countGoodTriplets(arr []int, a int, b int, c int) int {
    count := 0
    for i := 0; i < len(arr)-2; i++ {
        for j := i+1; j < len(arr)-1; j++ {
            for k := j+1; k < len(arr); k++ {
                if abs(arr[i] - arr[j]) <= a && abs(arr[j] - arr[k]) <= b && abs(arr[i] - arr[k]) <= c {
                    count++
                }
            }
        }
    }

    return count
}

// brute force
// time: O(n^3)
// space: O(1)
 
func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}