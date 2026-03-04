class Solution:
    def sortByBits(self, arr: List[int]) -> List[int]:
        # built-in func count set bits
        return sorted(arr, key=lambda x: (x.bit_count(), x)) # asc (set bits x, x)

# array, bit manipulation, sorting
# time: O(nlog(n))
# space: O(n)