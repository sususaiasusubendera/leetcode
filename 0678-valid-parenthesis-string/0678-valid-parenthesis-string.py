class Solution:
    def checkValidString(self, s: str) -> bool:
        min_balance = max_balance = 0
        for c in s:
            if c == '(':
                min_balance += 1
                max_balance += 1
            elif c == ')':
                min_balance -= 1
                max_balance -= 1
            else: # c == '*'
                min_balance -= 1
                max_balance += 1
            
            # not valid
            if max_balance < 0:
                return False
            
            # clip negative-balance for min_balance
            min_balance = max(min_balance, 0)
        
        return min_balance == 0

# greedy, string
# a beautiful and clean approach form la_castille
# time: O(n)
# space: O(1)