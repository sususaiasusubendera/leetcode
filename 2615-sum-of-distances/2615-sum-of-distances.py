class Solution:
    def distance(self, nums: List[int]) -> List[int]:
        ans = [0] * len(nums)

        m = {}
        for i, num in enumerate(nums):
            if num not in m:
                m[num] = []
            m[num].append(i)

        for k in m:
            if len(m[k]) > 1:
                arr = m[k]
                
                sum_pos, sum_neg = 0, 0
                count = 0
                for i in range(1, len(arr)):
                    sum_pos += arr[i]
                    count += 1
                
                for i in range(len(arr)):
                    ans[arr[i]] = (arr[i] * (i - count)) - sum_neg + sum_pos
                    if i < len(arr) - 1:
                        count -= 1
                        sum_neg += arr[i]
                        sum_pos -= arr[i + 1]
        
        return ans

# array, hash map, math, prefix sum
# time: O(n)
# space: O(n)

# tips
# go grab some paper and a pen :)