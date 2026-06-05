class Solution:
    def minimumCost(self, cost: List[int]) -> int:
        cost.sort(reverse=True)

        n = len(cost)
        idx = 0
        total = 0
        while idx <= n - 1:
            if idx <= n - 3:
                total += cost[idx] + cost[idx + 1]
                idx += 3
            elif idx == n - 2:
                total += cost[idx] + cost[idx + 1]
                idx += 2
            elif idx == n - 1:
                total += cost[idx]
                idx += 1
        
        return total

# array, greedy, sorting
# time: O(nlog(n))
# space: O(1)