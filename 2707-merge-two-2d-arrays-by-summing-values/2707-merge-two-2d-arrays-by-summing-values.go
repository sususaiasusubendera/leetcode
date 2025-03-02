func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    mapIdIndex:= map[int]int{} // ID -> index (result)
    var result [][]int
    p1, p2 := 0, 0
    for p1 < len(nums1) && p2 < len(nums2) {
        if nums1[p1][0] < nums2[p2][0] {
            if _, exists := mapIdIndex[nums1[p1][0]]; !exists {
                result = append(result, []int{nums1[p1][0], nums1[p1][1]})
                mapIdIndex[nums1[p1][0]] = len(result)-1
            } else {
                result[mapIdIndex[nums1[p1][0]]][1] += nums1[p1][1]
            }
            p1++
        } else if nums1[p1][0] > nums2[p2][0] {
            if _, exists := mapIdIndex[nums2[p2][0]]; !exists {
                result = append(result, []int{nums2[p2][0], nums2[p2][1]})
                mapIdIndex[nums2[p2][0]] = len(result)-1
            } else {
                result[mapIdIndex[nums2[p2][0]]][1] += nums2[p2][1]
            }
            p2++
        } else {
            result = append(result, []int{nums1[p1][0], nums1[p1][1]+nums2[p2][1]})
            p1++
            p2++
        }
    }

    // add the remaining element(s) from nums1
    for p1 < len(nums1) {
        result = append(result, []int{nums1[p1][0], nums1[p1][1]})
        p1++
    }

    // add the remaining element(s) from nums2
    for p2 < len(nums2) {
        result = append(result, []int{nums2[p2][0], nums2[p2][1]})
        p2++
    }

    return result
}

// one pass
// time: O(m + n)
// space: O(m + n)