class Solution:
    def maximizeSquareArea(self, m: int, n: int, hFences: List[int], vFences: List[int]) -> int:
        hFences, vFences = [1] + hFences + [m], [1] + vFences + [n]
        
        hFences.sort()
        vFences.sort()

        hSet = set()
        for i in range(0, len(hFences)-1):
            for j in range(i + 1, len(hFences)):
                s = hFences[j] - hFences[i]
                hSet.add(s)
        
        max_s = -1
        for i in range(0, len(vFences)-1):
            for j in range(i + 1, len(vFences)):
                s = vFences[j] - vFences[i]
                if s in hSet:
                    max_s = max(max_s, s)

        if max_s == -1:
            return -1
        
        MOD = 1_000_000_007

        return (max_s * max_s) % MOD

# array, hash map
# time: O(h^2 + v^2)
# space: O(h^2 + h + v)