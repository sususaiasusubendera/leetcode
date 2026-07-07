class Solution:
    def removeCoveredIntervals(self, intervals: List[List[int]]) -> int:
        if len(intervals) == 1:
            return 1

        intervals.sort(key=lambda x: (x[0], -x[1]))
        res = 1
        right = intervals[0][1]
        for start, end in intervals:
            if end <= right:
                continue
            else:
                res += 1
                right = end
        
        return res

# a beautiful approach from la_castille a.k.a eunice a.k.a ava taylor swift
# array, sorting
# time: O(nlog(n))
# space: O(1)