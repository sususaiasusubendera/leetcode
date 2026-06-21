class Solution:
    def maxIceCream(self, costs: List[int], coins: int) -> int:
        costs.sort()

        ice_cream, total_cost = 0, 0
        for i in range(len(costs)):
            if total_cost + costs[i] > coins:
                return ice_cream
            total_cost += costs[i]
            ice_cream += 1

        return ice_cream

# array, greedy
# time: O(nlog(n))
# space: O(1)