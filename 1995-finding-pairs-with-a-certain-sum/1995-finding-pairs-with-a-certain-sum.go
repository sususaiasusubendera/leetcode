type FindSumPairs struct {
    nums1 []int
    nums2 []int
    count map[int]int
}


func Constructor(nums1 []int, nums2 []int) FindSumPairs {
    count := make(map[int]int)
    for _, num := range nums2 {
        count[num]++
    }

    return FindSumPairs{
        nums1: nums1,
        nums2: nums2,
        count: count,
    }
}


func (this *FindSumPairs) Add(index int, val int)  {
    oldVal := this.nums2[index]
    this.count[oldVal]--
    this.nums2[index] += val
    this.count[this.nums2[index]]++
}


func (this *FindSumPairs) Count(tot int) int {
    result := 0
    for _, num := range this.nums1 {
        rest := tot - num
        result += this.count[rest]
    }

    return result
}

// time
// - Constructor: O(n2)
// - Add: O(1)
// - Count: O(n1)
// space: O(n1 + n2)

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */