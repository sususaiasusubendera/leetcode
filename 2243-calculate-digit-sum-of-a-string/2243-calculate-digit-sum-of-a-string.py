class Solution:
    def digitSum(self, s: str, k: int) -> str:
        while len(s) > k:
            temp = ""
            for i in range(0, len(s), k):
                val = 0
                for j in range(i, min(len(s), i + k)):
                    val += int(s[j])
                temp += str(val)
            s = temp
        return s

# simulation, string
# time: O(n) ? O(nlog(n))
# space: O(n)