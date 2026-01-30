class Solution:
    def minimumCost(self, source: str, target: str, original: List[str], changed: List[str], cost: List[int]) -> int:
        dist = [[inf] * 26 for _ in range(26)]
        for i in range(len(dist)):
            dist[i][i] = 0
        
        for i in range(len(original)):
            s, t, c = ord(original[i]) - ord('a'), ord(changed[i]) - ord('a'), cost[i] # source, target, cost
            dist[s][t] = min(dist[s][t], c)
        
        for k in range(26):
            for i in range(26):
                for j in range(26):
                    dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
        
        ans = 0
        for i in range(len(source)):
            if source[i] == target[i]:
                continue
            
            s, t = ord(source[i]) - ord('a'), ord(target[i]) - ord('a')
            if dist[s][t] == inf:
                return -1
            ans += dist[s][t]

        return ans

# array, floyd-warshall, graph, string
# time: O(n)
# space: O(1)