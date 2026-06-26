class Solution:
    def assignEdgeWeights(self, edges: List[List[int]]) -> int:
        MOD = 10**9 + 7

        n = len(edges) + 1
        adj = [[] for _ in range(n + 1)]
        for u, v in edges:
            adj[u].append(v)
            adj[v].append(u)

        max_depth = 0

        def dfs(node: int, parent: int, depth: int) -> None:
            nonlocal max_depth

            if depth > max_depth:
                max_depth = depth
            
            for neigh in adj[node]:
                if neigh != parent:
                    dfs(neigh, node, depth + 1)
            

        dfs(1, 0, 0)

        return pow(2, max_depth - 1, MOD)

# dfs, math, tree
# time: O(n)
# space: O(n)