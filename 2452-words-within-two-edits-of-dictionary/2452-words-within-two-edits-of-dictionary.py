class Solution:
    def twoEditWords(self, queries: List[str], dictionary: List[str]) -> List[str]:
        ans = []
        for q in queries:
            for d in dictionary:
                if q == d:
                    ans.append(q)
                    break
                
                edit = 0
                valid = True
                for i in range(len(q)):
                    if q[i] != d[i]:
                        edit += 1
                    if edit > 2:
                        valid = False
                        break
                
                if valid and edit <= 2:
                    ans.append(q)
                    break   
        
        return ans

# array, brute force, string
# time: O(qdl)
# space: O(q)