class Solution:
    def solveQueries(self, nums: List[int], queries: List[int]) -> List[int]:
        n = len(nums)
        ans = [-1] * len(queries)

        m = {}
        for i in range(n):
            if nums[i] not in m:
                m[nums[i]] = []
            m[nums[i]].append(i)

        for i, q in enumerate(queries):
            if len(m[nums[q]]) > 1:
                idx = -1
                curr_list = m[nums[q]]
                l, r = 0, len(curr_list) - 1  # left, right
                while l <= r:
                    mid = l + ((r - l) // 2)
                    if curr_list[mid] == q:
                        idx = mid
                        break
                    elif curr_list[mid] < q:
                        l = mid + 1
                    else:
                        r = mid - 1

                if idx == 0:
                    ans[i] = min(
                        curr_list[0] + (n - curr_list[len(curr_list) - 1]),
                        curr_list[1] - curr_list[0],
                    )
                elif idx == len(curr_list) - 1:
                    ans[i] = min(
                        curr_list[len(curr_list) - 1] - curr_list[len(curr_list) - 2],
                        (n - curr_list[len(curr_list) - 1]) + curr_list[0],
                    )
                else:
                    ans[i] = min(
                        curr_list[idx] - curr_list[idx - 1],
                        curr_list[idx + 1] - curr_list[idx],
                    )

        return ans

# array, binary search, hash map
# time: O(n + qlog(k))
# space: O(n + q)