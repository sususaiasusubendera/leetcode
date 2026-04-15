class Solution:
    def closestTarget(self, words: List[str], target: str, startIndex: int) -> int:
        n = len(words)
        found = False
        ld, rd = 0, 0

        # next
        i = startIndex
        while True:
            if words[i] == target:
                found = True
                break
            
            i = (i + 1) % n
            rd += 1
            if i == startIndex:
                break
        
        if not found:
            return -1

        # prev
        i = startIndex
        while True:
            if words[i] == target:
                break
            
            i = (i - 1 + n) % n
            ld += 1
            if i == startIndex:
                break 
        
        return min(rd, ld)

# array, string
# time: O(n)
# space: O(1)