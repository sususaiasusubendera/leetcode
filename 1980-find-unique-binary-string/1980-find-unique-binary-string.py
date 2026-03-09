class Solution:
    def findDifferentBinaryString(self, nums: List[str]) -> str:
        def dfs(curr):
            if len(curr) == n:
                if curr not in nums:
                    return curr
                return ""
            
            add_zero = dfs(curr + "0")
            if add_zero:
                return add_zero

            add_one = dfs(curr + "1")
            return add_one

        n = len(nums)
        nums = set(nums)
        return dfs("")

# array, backtracking, hash map, string
# time: O(2^n)
# space: O(n)