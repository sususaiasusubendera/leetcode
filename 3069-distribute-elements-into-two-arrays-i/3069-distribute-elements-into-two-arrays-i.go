func resultArray(nums []int) []int {
	arr1, arr2 := []int{}, []int{}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			arr1 = append(arr1, nums[i])
		} else if i == 1 {
			arr2 = append(arr2, nums[i])
		} else {
			if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
				arr1 = append(arr1, nums[i])
			} else {
				arr2 = append(arr2, nums[i])
			}
		}
	}

	return append(arr1, arr2...)
}

// array
// time: O(n)
// space: O(n)