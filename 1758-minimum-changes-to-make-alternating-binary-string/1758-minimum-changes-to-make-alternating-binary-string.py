class Solution:
    def minOperations(self, s: str) -> int:
        ops_zero_even = 0
        for i in range(0, len(s)):
            if i % 2 == 0:
                if s[i] == '1': ops_zero_even += 1
            else:
                if s[i] == '0': ops_zero_even += 1
                
        ops_zero_odd = 0
        for i in range(0, len(s)):
            if i % 2 == 0:
                if s[i] == '0': ops_zero_odd += 1
            else:
                if s[i] == '1': ops_zero_odd += 1
        
        return min(ops_zero_even, ops_zero_odd)

# string
# time: O(n)
# space: O(1)