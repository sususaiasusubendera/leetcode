class Solution:
    def getCommon(self, nums1: List[int], nums2: List[int]) -> int:
        m, n = len(nums1), len(nums2)
        p1, p2 = 0, 0
        while p1 < m and p2 < n:
            if nums1[p1] == nums2[p2]:
                return nums1[p1]
            elif nums1[p1] < nums2[p2]:
                p1 += 1
            else:
                p2 += 1
        
        return -1

# two pointers
# time: O(m + n)
# space: O(1)