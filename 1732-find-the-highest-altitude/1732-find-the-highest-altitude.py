class Solution:
    def largestAltitude(self, gain: List[int]) -> int:
        ans = 0
        altitude = 0
        for g in gain:
            altitude += g
            ans = max(ans, altitude)
        
        return ans

# array
# time: O(n)
# space: O(1)