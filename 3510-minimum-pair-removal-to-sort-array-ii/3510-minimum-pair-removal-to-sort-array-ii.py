class Solution:
    def minimumPairRemoval(self, nums: List[int]) -> int:
        count_decrease = 0
        ops = 0
        h = []
        merged = [False] * len(nums)
        head = Node(nums[0], 0)
        curr = head
        for i in range(1, len(nums)):
            new_node = Node(nums[i], i)
            curr.next = new_node
            new_node.prev = curr

            pair = Pair(curr, new_node, curr.val + new_node.val)
            heapq.heappush(h, (pair.cost, pair.first.left, id(pair), pair))

            if nums[i - 1] > nums[i]:
                count_decrease += 1

            curr = new_node

        while count_decrease > 0:
            _, _, _, p = heapq.heappop(h)
            first, second, cost = p.first, p.second, p.cost

            if (
                merged[first.left]
                or merged[second.left]
                or first.val + second.val != cost
            ):
                continue

            ops += 1
            if first.val > second.val:
                count_decrease -= 1

            prev_node = first.prev
            next_node = second.next
            first.next = next_node
            if next_node:
                next_node.prev = first

            if prev_node:
                # before - after
                if prev_node.val > first.val and prev_node.val <= cost:
                    count_decrease -= 1
                elif prev_node.val <= first.val and prev_node.val > cost:
                    count_decrease += 1

                pair = Pair(prev_node, first, prev_node.val + cost)
                heapq.heappush(h, (pair.cost, pair.first.left, id(pair), pair)) # tricked using id(pair)

            if next_node:
                # before - after
                if second.val > next_node.val and cost <= next_node.val:
                    count_decrease -= 1
                elif second.val <= next_node.val and cost > next_node.val:
                    count_decrease += 1

                pair = Pair(first, next_node, cost + next_node.val)
                heapq.heappush(h, (pair.cost, pair.first.left, id(pair), pair)) # tricked using id(pair)

            first.val = cost
            merged[second.left] = True

        return ops


# array, doubly-linked list, heap, linked list, simulation
# time: O(nlog(n))
# space: O(n)


class Node:
    def __init__(self, val, left):
        self.val = val
        self.left = left
        self.prev = None
        self.next = None


class Pair:
    def __init__(self, first, second, cost):
        self.first = first
        self.second = second
        self.cost = cost
