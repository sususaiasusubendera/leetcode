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