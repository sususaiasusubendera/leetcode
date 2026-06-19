class Solution:
    def angleClock(self, hour: int, minutes: int) -> float:
        minutesDeg = (minutes / 5) * 30
        hourDeg = ((hour + (minutes / 60)) % 12) * 30
        deg = abs(minutesDeg - hourDeg)
        return deg if deg < 180 else 360 - deg

# math
# time: O(1)
# space: O(1)