class Solution:
    def sortByBits(self, arr: List[int]) -> List[int]:
        return sorted(arr, key= lambda a: (a.bit_count(), a))

# solution from solutions (sound_wave9)
# notice me senpai!!