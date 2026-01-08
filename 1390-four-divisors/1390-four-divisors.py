class Solution:
    def sumFourDivisors(self, nums: List[int]) -> int:
        memo = {} # num -> count_div; memo[num] = -1 if invalid

        ans = 0
        for num in nums:
            if num == 1:
                continue
            
            if num in memo:
                if memo[num] != -1:
                    ans += memo[num]
                continue
                
            div = 2
            count_div = 2 # 1 and num included
            sum = 1 + num
            while div * div <= num:
                if num % div == 0:
                    if div == (num // div):
                        sum += div
                        count_div += 1
                    else:
                        sum += div + (num // div)
                        count_div += 2
                
                div += 1

                if count_div > 4:
                    break
            
            if count_div == 4:
                ans += sum
                memo[num] = sum
            else:
                memo[num] = -1
        
        return ans

# array, hash map, map
# time: O(n * sqrt(m))
# space: O(n)
