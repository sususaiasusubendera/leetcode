class Solution:
    def pivotArray(self, nums: List[int], pivot: int) -> List[int]:
        l = [] # elements less than pivot
        e = [] # elements equal to pivot
        g = [] # elements greater than pivot
        for num in nums:
            if num < pivot:
                l.append(num)
            elif num > pivot:
                g.append(num)
            else:
                e.append(num)
        
        return l + e + g

# array
# time: O(n)
# space: O(n)