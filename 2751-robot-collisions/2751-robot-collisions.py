class Solution:
    def survivedRobotsHealths(self, positions: List[int], healths: List[int], directions: str) -> List[int]:
        n = len(positions)

        # (robot's original idx, pos, health, dir)
        robots = [(i + 1, positions[i], healths[i], directions[i]) for i in range(n)]

        robots.sort(key=lambda x: x[1])

        stack = []
        for idx, pos, health, dir in robots:
            curr_health = health

            while stack and stack[-1][3] == 'R' and dir == 'L' and stack[-1][2] < curr_health:
                curr_health -= 1
                stack.pop()
            
            if stack and stack[-1][3] == 'R' and dir == 'L' and stack[-1][2] > curr_health:
                top = stack[-1]
                stack[-1] = (top[0], top[1], top[2] - 1, top[3])
            elif stack and stack[-1][3] == 'R' and dir == 'L' and stack[-1][2] == curr_health:
                stack.pop()
            else:
                stack.append((idx, pos, curr_health, dir))
        
        stack.sort(key=lambda x: x[0])

        return [robot[2] for robot in stack]

# array, sorting, stack
# time: O(nlog(n))
# space: O(n)