class Solution:
    def minimumDistance(self, nums: List[int]) -> int:
        n = len(nums)
        if n < 3:
            return -1

        m = {}
        x = 10**10
        found = False
        for i in range(n):
            if nums[i] not in m:
                m[nums[i]] = [i]
            else:
                m[nums[i]].append(i)
                nn = len(m[nums[i]])
                if nn >= 3:
                    a, b, c = m[nums[i]][nn - 3], m[nums[i]][nn - 2], m[nums[i]][nn - 1]
                    x = min(x, abs(a - b) + abs(b - c) + abs(c - a))
                    found = True
        
        if found:
            return x
        return -1

# array, hash map
# time: O(n)
# space: O(n)