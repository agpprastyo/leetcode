// Last updated: 2/11/2026, 8:30:32 PM
func searchInsert(nums []int, target int) int {
    left, right := 0, len (nums)-1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return left
}