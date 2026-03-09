class Solution:
    def minFlips(self, s: str) -> int:
        n = len(s)

        s2 = s + s

        o = '10' * n # alternating string start with 1 size 2 * n
        z = '01' * n # alternating string start with 0 size 2 * n

        ans = 10**10
        on, zn = 0, 0
        left, right = 0, 0
        while right < n * 2:
            if s2[right] != o[right]:
                on += 1
            if s2[right] != z[right]:
                zn += 1

            if right - left + 1 > n:
                if s2[left] != o[left]:
                    on -= 1
                if s2[left] != z[left]:
                    zn -= 1
                left += 1

                ans = min(ans, min(on, zn))
            
            right += 1

        return ans

# sliding window, string
# time: O(n)
# space: O(n)