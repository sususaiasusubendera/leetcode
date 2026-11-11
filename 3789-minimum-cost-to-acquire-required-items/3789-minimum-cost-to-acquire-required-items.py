class Solution:
    def minimumCost(self, cost1: int, cost2: int, costBoth: int, need1: int, need2: int) -> int:
        if not need1 and not need2:
            return 0
        
        a = b = c = d = 0

        # costBoth (type 1 and type 2) all in
        c = max(need1, need2) * costBoth

        # type 1 all in
        a = need1 * cost1

        # type 2 all in
        b = need2 * cost2

        # costBoth + type 1 or type 2 (greedy)
        min_one = True
        if need1 > need2:
            min_one = False
        d = min(need1, need2) * costBoth
        if min_one:
            d += (need2 - need1) * cost2
        else:
            d += (need1 - need2) * cost1
        

        return min(a + b, c, d)

# greedy, math
# time: O(1)
# space: O(1)