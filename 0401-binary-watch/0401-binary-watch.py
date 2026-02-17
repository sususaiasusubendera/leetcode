class Solution:
    def readBinaryWatch(self, turnedOn: int) -> List[str]:
        ans = []
        def dfs(pos, count, h, m):
            if h > 11 or m > 59: return

            if pos == 10:
                if count == turnedOn:
                    ans.append(f"{h}:{m:02d}")
                return

            # off
            dfs(pos + 1, count, h, m)

            # on
            if pos < 4:
                dfs(pos + 1, count + 1, h + (1 << pos), m)
            else:
                dfs(pos + 1, count + 1, h, m + (1 << (pos - 4)))

        
        dfs(0, 0, 0, 0)
        return ans