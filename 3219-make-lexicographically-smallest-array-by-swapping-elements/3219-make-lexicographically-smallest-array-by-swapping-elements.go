func lexicographicallySmallestArray(nums []int, limit int) []int {
    // copy and sort it
    numsSorted := make([]int, len(nums))
    copy(numsSorted, nums)
    sort.Ints(numsSorted)

    // create a map to group elements based on "limit"
    currGroup := 0
    numToGroup := map[int]int{}
    groupToList := map[int][]int{}
    for i := 0; i < len(numsSorted); i++ {
        // make a new group if it's a first element or have difference with prev element > limit 
        if i == 0 || numsSorted[i] - numsSorted[i-1] > limit {
            currGroup++
        }

        // insert the element into the group
        numToGroup[numsSorted[i]] = currGroup
        groupToList[currGroup] = append(groupToList[currGroup], numsSorted[i])
    }

    // go back to nums, and reorder the elements based on the sorted groups.
    for i := 0; i < len(nums); i++ {
        group := numToGroup[nums[i]] // current group
        nums[i] = groupToList[group][0]
        groupToList[group] = groupToList[group][1:] // delete the element that has been used
    }

    return nums
}

// sorting + grouping approach
// time: O(nlogn)
// space: O(n)