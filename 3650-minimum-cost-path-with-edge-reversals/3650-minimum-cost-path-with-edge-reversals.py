class Solution:
    def minCost(self, n: int, edges: List[List[int]]) -> int:
        adj = [[] for _ in range(n)]
        for u, v, w in edges:
            adj[u].append((v, w))
            adj[v].append((u, 2 * w))
        
        d = [inf] * n
        d[0] = 0
        visited = [False] * n
        h = [(0, 0)] # (distance, node)
        while h:
            curr_d, node = heapq.heappop(h)

            if node == n - 1:
                return curr_d
            
            if visited[node]:
                continue
            
            visited[node] = True
            for v, w in adj[node]:
                new_d = curr_d + w
                if new_d < d[v]:
                    d[v] = new_d
                    heapq.heappush(h, (new_d, v))
        
        return -1

# dijkstra, graph, heap
# time: O(nlog(n))
# space: O(n)