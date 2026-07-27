class Solution:
    def uniqueXorTriplets(self, nums: List[int]) -> int:
        MAX = 2047

        a, b, c = [False] * (MAX + 1), [False] * (MAX + 1), [False] * (MAX + 1)

        for n in nums:
            a[n] = True
        
        for i in range(MAX):
            if not a[i]:
                continue

            for n in nums:
                b[i ^ n] = True
        
        for i in range(MAX):
            if not b[i]:
                continue
            
            for n in nums:
                c[i ^ n] = True
        
        ans = 0
        for ok in c:
            if ok:
                ans += 1
        
        return ans

# array, brute force + optimization
# time: O(MAX * n)
# space: O(MAX)