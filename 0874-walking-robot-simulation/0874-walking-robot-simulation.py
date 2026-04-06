class Solution:
    def robotSim(self, commands: List[int], obstacles: List[List[int]]) -> int:
        o = set((x, y) for x, y in obstacles)

        dirs = [
            (0, 1), # north
            (1, 0), # east
            (0, -1), # south
            (-1, 0), # west
        ]

        ans = 0
        face = 0 # face: 0 (north), 1 (east), 2 (south), 3 (west)
        p = (0, 0) # p (x, y)
        for c in commands:
            if 1 <= c <= 9:
                for i in range(c):
                    nx, ny = p[0] + dirs[face][0], p[1] + dirs[face][1]
                    if (nx, ny) not in o:
                        p = (nx, ny)
                ans = max(ans, p[0]**2 + p[1]**2)
            elif c == -1 or c == -2:
                face = turn(c, face)
        
        return ans

def turn(c, f):
    if c == -1:
        f = (f + 1) % 4
    elif c == -2:
        f = (f - 1 + 4) % 4
    return f

# array, hash map
# time: O(n)
# space: O(m)