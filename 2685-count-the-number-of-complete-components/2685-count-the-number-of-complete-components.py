class Solution:
    def countCompleteComponents(self, n: int, edges: List[List[int]]) -> int:
        adj = [[] for _ in range(n)]
        for i in range(n):
            adj[i].append(i)
        
        for u, v in edges:
            adj[u].append(v)
            adj[v].append(u)
        
        m = {}
        for i in range(len(adj)):
            adj[i].sort()
            m[tuple(adj[i])] = m.get(tuple(adj[i]), 0) + 1

        ans = 0
        for k in m:
            if len(k) == m[k]:
                ans += 1
        
        return ans

# adjacency list approach, graph
# time: O(mlog(n))
# space: O(n + m)
