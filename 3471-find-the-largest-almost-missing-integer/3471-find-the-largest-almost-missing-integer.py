class Solution:
    def largestInteger(self, nums: List[int], k: int) -> int:
        freq = [0] * 51
        for i in range(len(nums) - k + 1):
            seen = set()
            for j in range(i, i + k):
                if nums[j] not in seen:
                    seen.add(nums[j])
                    freq[nums[j]] += 1
        
        ans = -1
        for num in range(len(freq)):
            if freq[num] == 1:
                ans = max(ans, num)
        
        return ans

# array, hash map
# time: O(nk)
# space: O(k)