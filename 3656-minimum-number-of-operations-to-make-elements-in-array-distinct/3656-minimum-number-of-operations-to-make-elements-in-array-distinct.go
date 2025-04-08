func minimumOperations(nums []int) int {
  p1, p2 := 0, 0 // pointer
  ops := 0
  numSet := map[int]bool{}
  for p2 < len(nums) && p1 < len(nums) {
    if !numSet[nums[p2]] {
        numSet[nums[p2]] = true
        p2++
        continue
    }

    ops++
    p1 += 3
    p2 = p1
    numSet = map[int]bool{}
  }

  return ops
}

// greedy two pointer
// time: O(n)
// space: O(n)