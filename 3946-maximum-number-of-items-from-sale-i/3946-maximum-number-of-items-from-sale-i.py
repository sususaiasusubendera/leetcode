class Solution:
    def maximumSaleItems(self, items: List[List[int]], budget: int) -> int:
        dp = [0] * (budget + 1)
        min_price = min(price for _, price in items)
        for factor, price in items:
            count = sum(
                1 for other_factor, _ in items if other_factor % factor == 0
            )
            for j in range(budget, price - 1, -1):
                if dp[j - price] + count > dp[j]:
                    dp[j] = dp[j - price] + count
        return max(dp[j] + (budget - j) // min_price for j in range(budget + 1))

# editorial (dp 0-1 knapsack)
# notice me senpai