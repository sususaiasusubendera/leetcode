class Solution:
    def maxNumberOfFamilies(self, n: int, reservedSeats: List[List[int]]) -> int:
        ans = 2 * n

        reservedSeats.sort()

        idx = 0
        while idx < len(reservedSeats):
            curr_row = reservedSeats[idx][0]
            left, mid_left, mid_right, right = False, False, False, False
            while idx < len(reservedSeats) and reservedSeats[idx][0] == curr_row:
                seat = reservedSeats[idx][1]

                if seat == 2 or seat == 3:
                    left = True
                elif seat == 4 or seat == 5:
                    mid_left = True
                elif seat == 6 or seat == 7:
                    mid_right = True
                elif seat == 8 or seat == 9:
                    right = True
                
                idx += 1
            
            if (left and mid_right) or (mid_left and mid_right) or (mid_left and right):
                ans -= 2
            elif left or mid_left or mid_right or right:
                ans -= 1
        
        return ans

# array, greedy, sorting
# time: O(nlog(n))
# space: O(1)