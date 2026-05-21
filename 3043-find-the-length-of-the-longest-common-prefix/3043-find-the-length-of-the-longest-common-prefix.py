class Solution:
    def longestCommonPrefix(self, arr1: List[int], arr2: List[int]) -> int:
        ans = 0
        prefixes = set()

        for v in arr1:
            while v > 0 and v not in prefixes:
                prefixes.add(v)
                v //= 10
        
        for v in arr2:
            while v > 0 and v not in prefixes:
                v //= 10
            
            if v == 0: continue

            ans = max(ans, len(str(v)))

        return ans

# array, hash map
# time: O(nlog(d) + mlog(d))
# space: O(nlog(d))

# #1 2026/05/21