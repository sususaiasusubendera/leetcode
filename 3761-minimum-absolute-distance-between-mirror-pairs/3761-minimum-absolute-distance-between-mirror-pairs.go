func minMirrorPairDistance(nums []int) int {
    prev := map[int]int{}
    ans := 1_000_000_007
    for i, num := range nums {
        if _, ok := prev[num]; ok {
            ans = min(ans, i - prev[num])
        }
        
        prev[reverse(num)] = i
    }

    if ans == 1_000_000_007 {
        return -1
    }
    return ans
}

func reverse(n int) int {
	res := 0
	for n > 0 {
		d := n % 10
		res = (res * 10) + d
        n /= 10
	}
    return res
}

// array, hash map
// time: O(n)
// space: O(n)