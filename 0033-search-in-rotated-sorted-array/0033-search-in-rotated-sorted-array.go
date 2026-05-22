func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
    for left <= right {
        mid := left + ((right - left) / 2)
        if nums[mid] == target {
            return mid
        } 
        
        if nums[left] > nums[right] { // left rotated
            if nums[left] <= nums[mid] { // left side is sorted, but right side is not
                if nums[left] <= target && target < nums[mid] {
                    right = mid - 1
                } else {
                    left = mid + 1
                }
            } else { // right side is sorted, but left side is not
                if nums[mid] < target && target <= nums[right] {
                    left = mid + 1
                } else {
                    right = mid - 1
                }
            }
        } else { // normal asc order
            if nums[mid] < target {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }

    return -1
}

// array, binary search
// time: O(log(n))
// space: O(1)