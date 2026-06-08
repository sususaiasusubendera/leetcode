func pivotArray(nums []int, pivot int) []int {
    less := []int{} // elements less than pivot
    equal := []int{} // element equal to pivot
    greater := []int{} // elements greater than pivot
    for _, num := range nums {
        if num < pivot {
            less = append(less, num)
        } else if num > pivot {
            greater = append(greater, num)
        } else {
            equal = append(equal, num)
        }
    }

    ans := make([]int, 0, len(nums)) // pre-allocate for efficiency
    ans = append(ans, less...)
    ans = append(ans, equal...)
    ans = append(ans, greater...)

    return ans
}

// array
// time: O(n)
// space: O(n)