class Solution:
    def remainingMethods(self, n: int, k: int, invocations: List[List[int]]) -> List[int]:
        adj = [[] for _ in range(n)]
        for u, v in invocations:
            adj[u].append(v)

        queue = deque([k])
        sus = [False] * n
        sus[k] = True
        while len(queue) > 0:
            u = queue.popleft()
            for v in adj[u]:
                if not sus[v]:
                    sus[v] = True
                    queue.append(v)
        
        for u, v in invocations:
            if not sus[u] and sus[v]:
                res = [0] * n
                for i in range(n):
                    res[i] = i
                return res
        
        res = []
        for i in range(n):
            if not sus[i]:
                res.append(i)
        
        return res

# bfs, graph
# time: O(m + n)
# space: O(m + n)