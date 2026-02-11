// Last updated: 2/11/2026, 8:30:45 PM
func maxArea(height []int) int {
    left, right := 0, len(height)-1
    maxArea :=  0


    for left < right {
        width := right - left
        h := min(height[left], height[right])

        area := width * h
        if area > maxArea {
            maxArea = area
        }

        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return maxArea
}