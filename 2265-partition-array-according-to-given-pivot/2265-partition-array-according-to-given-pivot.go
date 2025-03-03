func pivotArray(nums []int, pivot int) []int {
    var leftSlices []int
    var midSlices []int
    var rightSlices []int
    for i := 0; i < len(nums); i++ {
        if nums[i] < pivot {
            leftSlices = append(leftSlices, nums[i])
            continue
        } else if nums[i] > pivot {
            rightSlices = append(rightSlices, nums[i])
        } else {
            midSlices = append(midSlices, nums[i])
        }
    }

    var result []int
    for i := 0; i < len(leftSlices); i++ {
        result = append(result, leftSlices[i])
    }

    for i := 0; i < len(midSlices); i++ {
        result = append(result, midSlices[i])
    }    

    for i := 0; i < len(rightSlices); i++ {
        result = append(result, rightSlices[i])
    }

    return result
}

// brute force
// time: O(n)
// space: O(n)