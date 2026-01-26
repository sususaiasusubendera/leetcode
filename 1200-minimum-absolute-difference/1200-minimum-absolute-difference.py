class Solution:
    def minimumAbsDifference(self, arr: List[int]) -> List[List[int]]:
        arr.sort()

        min_diff = sys.maxsize
        for i in range(1, len(arr)):
            min_diff = min(min_diff, arr[i] - arr[i - 1])
        
        ans = []
        for i in range(1, len(arr)):
            diff = arr[i] - arr[i - 1]
            if diff == min_diff:
                ans.append([arr[i - 1], arr[i]])
        
        return ans

# array, sorting
# time: O(nlog(n))
# space: O(n)