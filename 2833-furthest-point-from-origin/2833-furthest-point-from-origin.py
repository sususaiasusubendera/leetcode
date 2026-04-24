class Solution:
    def furthestDistanceFromOrigin(self, moves: str) -> int:
        r, l, x = 0, 0, 0
        for m in moves:
            if m == 'R':
                r += 1
            elif m == 'L':
                l += 1
            else:
                x += 1
        
        return max(r, l) - min(r, l) + x

# counting, string
# time: O(n)
# space: O(1)