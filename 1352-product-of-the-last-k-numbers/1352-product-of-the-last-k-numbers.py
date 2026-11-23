class ProductOfNumbers:
    def __init__(self):
        self.pre = [1]

    def add(self, num: int) -> None:
        if not num:
            self.pre = [1]
        else:
            n = len(self.pre)
            self.pre.append(self.pre[n - 1] * num)

    def getProduct(self, k: int) -> int:
        n = len(self.pre)
        return 0 if k >= n else self.pre[n - 1] // self.pre[n - k - 1]

# array, prefix sum
# time: O(1)
# space: O(n)
        
# Your ProductOfNumbers object will be instantiated and called as such:
# obj = ProductOfNumbers()
# obj.add(num)
# param_2 = obj.getProduct(k)