func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
    d := 0 // d^2 = dx^2 + dy^2
    if xCenter < x1 || xCenter > x2 {
        d += min((x1 - xCenter) * (x1 - xCenter), (x2 - xCenter) * (x2 - xCenter))
    }
    if yCenter < y1 || yCenter > y2 {
        d += min((y1 - yCenter) * (y1 - yCenter), (y2 - yCenter) * (y2 - yCenter))  
    }
    
    return d <= radius * radius
}

// geometry, math
// time: O(1)
// space: O(1)