99 ms --> 26 ms
### Before (99 ms)
```go
func threeSum(nums []int) [][]int {
    nums = quickSort(nums)
    var result [][]int
    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        left, right := i+1, len(nums)-1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum == 0 {
                result = append(result, []int{nums[i], nums[left], nums[right]})
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                for left < right && nums[right] == nums[right-1] {
                    right--
                }
                left++
                right--
            } else if sum >= 0 {
                right--
            } else {
                left++
            }
        }
    }
    return result
}

// quick sorting
func quickSort(a []int) []int {
    if len(a) <= 1 {
        return a
    }

    pivotIndex := len(a) / 2
    pivot := a[pivotIndex]
    left, right := []int{}, []int{}
    for i := range a {
        if i == pivotIndex {
            continue
        }
        if a[i] < pivot {
            left = append(left, a[i])
        } else {
            right = append(right, a[i])
        }
    }

    return append(append(quickSort(left), pivot), quickSort(right)...)
}
```

### After (26 ms)
```go
func threeSum(nums []int) [][]int {
    // sort the 'nums'
    quickSort(nums, 0, len(nums)-1)

    // nums[0] positive --> all elements of nums are positive --> no sum result == 0
    if nums[0] >= 1 {
        return [][]int{}
    }

    result := [][]int{}
    for i := 0; i < len(nums)-2; i++ {
        // check duplicate
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        left, right := i+1, len(nums)-1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum == 0 {
                result = append(result, []int{nums[i], nums[left], nums[right]})
                for left < right && nums[left] == nums[left+1] { // check duplicate
                    left++
                }
                for left < right && nums[right] == nums[right-1] { // check duplicate
                    right--
                }
                left++
                right--
            } else if sum >= 0 {
                right--
            } else {
                left++
            }
        }
    }
    return result
}

// quick sorting
func quickSort(a []int, low, high int) { // more efficient quick sort (in-place quick sort, without creating new slices)
    if low < high { // process only when 'element of a' > 1
        p := partition(a, low, high)
        quickSort(a, low, p-1)
        quickSort(a, p+1, high)
    }
}

// support function for in-place quick sort
// input: []int, int, int --> output int (new pivot index)
func partition(a []int, low, high int) int {
    pivot := a[high]
    i := low - 1
    for j := low; j < high; j++ {
        if a[j] <= pivot {
            i++
            a[i], a[j] = a[j], a[i]
        }
    }
    a[i+1], a[high] = a[high], a[i+1]
    return i + 1
}
```
