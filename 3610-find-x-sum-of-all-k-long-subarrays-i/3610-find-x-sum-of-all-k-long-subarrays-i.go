func findXSum(nums []int, k int, x int) []int {
	ans := []int{}
	data := &Data{
        nums: []int{},
        freq: map[int]int{},
    }

    left := 0
    for right := 0; right < len(nums); right++ {
        if _, exists := data.freq[nums[right]]; !exists {
            data.nums = append(data.nums, nums[right])
        }
        data.freq[nums[right]]++

        if right - left + 1 >= k {
            sort.Slice(data.nums, func(i, j int) bool {
                if data.freq[data.nums[i]] == data.freq[data.nums[j]] {
                    return data.nums[i] > data.nums[j]
                }

                return data.freq[data.nums[i]] > data.freq[data.nums[j]]
            })

            sum := 0
            for i := 0; i < x && i < len(data.nums); i++ {
                sum += data.nums[i] * data.freq[data.nums[i]] // num * freq of num
            }
            ans = append(ans, sum)

            data.freq[nums[left]]--
            if data.freq[nums[left]] <= 0 {
                for i, num := range data.nums {
                    if nums[left] == num {
                        data.nums = append(data.nums[:i], data.nums[i+1:]...)
                        break
                    }
                }
                delete(data.freq, nums[left])
            }

            left++
        }
    }

    return ans
}

type Data struct {
    nums []int
    freq map[int]int
}

// array, hash map, sliding window
// time: O((n^2)log(n))
// space: O(n)