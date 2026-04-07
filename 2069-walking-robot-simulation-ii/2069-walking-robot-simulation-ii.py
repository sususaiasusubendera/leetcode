class Robot:
    dirs = ["North", "East", "South", "West"]
    steps = [(0, 1), (1, 0), (0, -1), (-1, 0)]

    def __init__(self, width: int, height: int):
        self.w = width
        self.h = height
        self.d = 1
        self.p = [0, 0]

    def step(self, num: int) -> None:
        cycle = 2 * (self.w + self.h) - 4
        num = num % cycle

        # edge case
        if num == 0:
            if self.p[0] == 0 and self.p[1] == 0:
                self.d = 2 # south
                return

        for i in range(num):
            nx, ny = self.p[0] + self.steps[self.d][0], self.p[1] + self.steps[self.d][1]
            if nx < 0 or nx >= self.w or ny < 0 or ny >= self.h:
                self.d = (self.d - 1 + 4) % 4
                nx, ny = self.p[0] + self.steps[self.d][0], self.p[1] + self.steps[self.d][1]
            self.p[0], self.p[1] = nx, ny

    def getPos(self) -> List[int]:
        return self.p

    def getDir(self) -> str:
        return self.dirs[self.d]
        
# design, simulation
# time: O(n)
# space: O(1)

# Your Robot object will be instantiated and called as such:
# obj = Robot(width, height)
# obj.step(num)
# param_2 = obj.getPos()
# param_3 = obj.getDir()