class Solution:
    def maximizeSquareHoleArea(self, n: int, m: int, hBars: List[int], vBars: List[int]) -> int:
        hBars.sort()
        vBars.sort()

        h_max, v_max = 1, 1
        h, v = 1, 1
        if len(hBars) > 1:
            for i in range(1, len(hBars)):
                if hBars[i] == hBars[i-1] + 1:
                    h += 1
                    h_max = max(h_max, h)
                else:
                    h = 1
        if len(vBars) > 1:
            for i in range(1, len(vBars)):
                if vBars[i] == vBars[i-1] + 1:
                    v += 1
                    v_max = max(v_max, v)
                else:
                    v = 1
        
        s = min(h_max, v_max) + 1
        return s**2

# array, sorting
# time:(O(nlogn + mlogm))
# space: O(1)